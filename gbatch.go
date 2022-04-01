package gbatch

func New[T any](s []T, size int) Batch[T] {
    return Batch[T]{
        s:    s,
        size: size,
    }
}

type Batch[T any] struct {
    s    []T
    size int
    curr int
}

func (b *Batch[T]) Batch() []T {
    btch := b.Curr()
    b.Next()

    return btch
}

func (b *Batch[T]) Curr() []T {
    l := len(b.s)
    if b.curr >= l {
        return nil
    }

    end := b.curr + b.size
    if end > l {
        end = l
    }

    return b.s[b.curr:end]
}

func (b *Batch[T]) Next() bool {
    b.curr += b.size
    return b.curr < len(b.s)
}
