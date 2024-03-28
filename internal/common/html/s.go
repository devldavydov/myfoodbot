package html

type S struct {
	val string
}

var _ IELement = (*S)(nil)

func NewS(val string) *S {
	return &S{val: val}
}

func (r *S) Build() string {
	return r.val
}
