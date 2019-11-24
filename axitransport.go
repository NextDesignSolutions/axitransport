package axitransport

type AxiCacheAttributes struct {
	Bufferable  bool
	Modifiable  bool
	Read_alloc  bool
	Write_alloc bool
}

const (
	BUFFERABLE    = true
	UNBUFFERABLE  = false
	MODIFIABLE    = true
	UNMODIFIABLE  = false
	READ_ALLOC    = true
	READ_UNALLOC  = false
	WRITE_ALLOC   = true
	WRITE_UNALLOC = true
	INCR_MODE     = true
	FIXED_MODE    = false
	LTTILE_ENDIAN = false
	BIG_ENDIAN    = true
)

type AxiTransport interface {
	SetCacheAttributes(bufferable bool, modifiable bool, read_alloc bool, write_alloc bool) error
	SetIncrMode(mode bool) error
	SetEndianess(endian bool) error
	ReadAxi(addr uint64, size int) (data []byte, err error)
	WriteAxi(addr uint64, data []byte) (err error)
	ReadAxi32(addr uint64) (value uint32, err error)
	WriteAxi32(addr uint64, value uint32) (err error)
	ReadAxi64(addr uint64) (value uint64, err error)
	WriteAxi64(addr uint64, value uint64) (err error)
}
