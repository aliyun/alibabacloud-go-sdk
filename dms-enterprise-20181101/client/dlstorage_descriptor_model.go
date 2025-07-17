// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLStorageDescriptor interface {
	dara.Model
	String() string
	GoString() string
	SetBucketCols(v []*string) *DLStorageDescriptor
	GetBucketCols() []*string
	SetColumns(v []*DLColumn) *DLStorageDescriptor
	GetColumns() []*DLColumn
	SetInputFormat(v string) *DLStorageDescriptor
	GetInputFormat() *string
	SetIsCompressed(v bool) *DLStorageDescriptor
	GetIsCompressed() *bool
	SetLocation(v string) *DLStorageDescriptor
	GetLocation() *string
	SetNumBuckets(v int32) *DLStorageDescriptor
	GetNumBuckets() *int32
	SetOriginalColumns(v []*DLColumn) *DLStorageDescriptor
	GetOriginalColumns() []*DLColumn
	SetOutputFormat(v string) *DLStorageDescriptor
	GetOutputFormat() *string
	SetParameters(v map[string]interface{}) *DLStorageDescriptor
	GetParameters() map[string]interface{}
	SetSerdeInfo(v *DLSerdeInfo) *DLStorageDescriptor
	GetSerdeInfo() *DLSerdeInfo
	SetSkewedInfo(v *DLSkewedInfo) *DLStorageDescriptor
	GetSkewedInfo() *DLSkewedInfo
	SetSortCols(v []*DLOrder) *DLStorageDescriptor
	GetSortCols() []*DLOrder
}

type DLStorageDescriptor struct {
	BucketCols      []*string              `json:"BucketCols,omitempty" xml:"BucketCols,omitempty" type:"Repeated"`
	Columns         []*DLColumn            `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	InputFormat     *string                `json:"InputFormat,omitempty" xml:"InputFormat,omitempty"`
	IsCompressed    *bool                  `json:"IsCompressed,omitempty" xml:"IsCompressed,omitempty"`
	Location        *string                `json:"Location,omitempty" xml:"Location,omitempty"`
	NumBuckets      *int32                 `json:"NumBuckets,omitempty" xml:"NumBuckets,omitempty"`
	OriginalColumns []*DLColumn            `json:"OriginalColumns,omitempty" xml:"OriginalColumns,omitempty" type:"Repeated"`
	OutputFormat    *string                `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	Parameters      map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SerdeInfo       *DLSerdeInfo           `json:"SerdeInfo,omitempty" xml:"SerdeInfo,omitempty"`
	SkewedInfo      *DLSkewedInfo          `json:"SkewedInfo,omitempty" xml:"SkewedInfo,omitempty"`
	SortCols        []*DLOrder             `json:"SortCols,omitempty" xml:"SortCols,omitempty" type:"Repeated"`
}

func (s DLStorageDescriptor) String() string {
	return dara.Prettify(s)
}

func (s DLStorageDescriptor) GoString() string {
	return s.String()
}

func (s *DLStorageDescriptor) GetBucketCols() []*string {
	return s.BucketCols
}

func (s *DLStorageDescriptor) GetColumns() []*DLColumn {
	return s.Columns
}

func (s *DLStorageDescriptor) GetInputFormat() *string {
	return s.InputFormat
}

func (s *DLStorageDescriptor) GetIsCompressed() *bool {
	return s.IsCompressed
}

func (s *DLStorageDescriptor) GetLocation() *string {
	return s.Location
}

func (s *DLStorageDescriptor) GetNumBuckets() *int32 {
	return s.NumBuckets
}

func (s *DLStorageDescriptor) GetOriginalColumns() []*DLColumn {
	return s.OriginalColumns
}

func (s *DLStorageDescriptor) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *DLStorageDescriptor) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *DLStorageDescriptor) GetSerdeInfo() *DLSerdeInfo {
	return s.SerdeInfo
}

func (s *DLStorageDescriptor) GetSkewedInfo() *DLSkewedInfo {
	return s.SkewedInfo
}

func (s *DLStorageDescriptor) GetSortCols() []*DLOrder {
	return s.SortCols
}

func (s *DLStorageDescriptor) SetBucketCols(v []*string) *DLStorageDescriptor {
	s.BucketCols = v
	return s
}

func (s *DLStorageDescriptor) SetColumns(v []*DLColumn) *DLStorageDescriptor {
	s.Columns = v
	return s
}

func (s *DLStorageDescriptor) SetInputFormat(v string) *DLStorageDescriptor {
	s.InputFormat = &v
	return s
}

func (s *DLStorageDescriptor) SetIsCompressed(v bool) *DLStorageDescriptor {
	s.IsCompressed = &v
	return s
}

func (s *DLStorageDescriptor) SetLocation(v string) *DLStorageDescriptor {
	s.Location = &v
	return s
}

func (s *DLStorageDescriptor) SetNumBuckets(v int32) *DLStorageDescriptor {
	s.NumBuckets = &v
	return s
}

func (s *DLStorageDescriptor) SetOriginalColumns(v []*DLColumn) *DLStorageDescriptor {
	s.OriginalColumns = v
	return s
}

func (s *DLStorageDescriptor) SetOutputFormat(v string) *DLStorageDescriptor {
	s.OutputFormat = &v
	return s
}

func (s *DLStorageDescriptor) SetParameters(v map[string]interface{}) *DLStorageDescriptor {
	s.Parameters = v
	return s
}

func (s *DLStorageDescriptor) SetSerdeInfo(v *DLSerdeInfo) *DLStorageDescriptor {
	s.SerdeInfo = v
	return s
}

func (s *DLStorageDescriptor) SetSkewedInfo(v *DLSkewedInfo) *DLStorageDescriptor {
	s.SkewedInfo = v
	return s
}

func (s *DLStorageDescriptor) SetSortCols(v []*DLOrder) *DLStorageDescriptor {
	s.SortCols = v
	return s
}

func (s *DLStorageDescriptor) Validate() error {
	return dara.Validate(s)
}
