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
	CompressionLevel *int32 `json:"CompressionLevel,omitempty" xml:"CompressionLevel,omitempty"`
	// example:
	//
	// draco
	LibraryName      *string `json:"LibraryName,omitempty" xml:"LibraryName,omitempty"`
	QuantizationBits *int32  `json:"QuantizationBits,omitempty" xml:"QuantizationBits,omitempty"`
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
