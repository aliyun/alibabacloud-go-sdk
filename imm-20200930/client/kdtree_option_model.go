// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKdtreeOption interface {
	dara.Model
	String() string
	GoString() string
	SetCompressionLevel(v int32) *KdtreeOption
	GetCompressionLevel() *int32
	SetLibraryName(v string) *KdtreeOption
	GetLibraryName() *string
	SetQuantizationBits(v int32) *KdtreeOption
	GetQuantizationBits() *int32
}

type KdtreeOption struct {
	// The compression level. Valid values: 0 to 10. A greater value specifies a higher compression ratio and ensures better detail effects.
	//
	// example:
	//
	// 1
	CompressionLevel *int32 `json:"CompressionLevel,omitempty" xml:"CompressionLevel,omitempty"`
	// The name of the library supported by a k-d tree. Set the value to draco. Default value: draco.
	//
	// example:
	//
	// draco
	LibraryName *string `json:"LibraryName,omitempty" xml:"LibraryName,omitempty"`
	// The number of bits for quantization. Valid values: 0 to 31. A greater value ensures that more details are retained. A value of 0 specifies that vertex compression is not performed.
	//
	// example:
	//
	// 1
	QuantizationBits *int32 `json:"QuantizationBits,omitempty" xml:"QuantizationBits,omitempty"`
}

func (s KdtreeOption) String() string {
	return dara.Prettify(s)
}

func (s KdtreeOption) GoString() string {
	return s.String()
}

func (s *KdtreeOption) GetCompressionLevel() *int32 {
	return s.CompressionLevel
}

func (s *KdtreeOption) GetLibraryName() *string {
	return s.LibraryName
}

func (s *KdtreeOption) GetQuantizationBits() *int32 {
	return s.QuantizationBits
}

func (s *KdtreeOption) SetCompressionLevel(v int32) *KdtreeOption {
	s.CompressionLevel = &v
	return s
}

func (s *KdtreeOption) SetLibraryName(v string) *KdtreeOption {
	s.LibraryName = &v
	return s
}

func (s *KdtreeOption) SetQuantizationBits(v int32) *KdtreeOption {
	s.QuantizationBits = &v
	return s
}

func (s *KdtreeOption) Validate() error {
	return dara.Validate(s)
}
