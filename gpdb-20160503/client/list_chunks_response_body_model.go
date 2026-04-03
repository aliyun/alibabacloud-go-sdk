// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChunksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChunks(v *ListChunksResponseBodyChunks) *ListChunksResponseBody
	GetChunks() *ListChunksResponseBodyChunks
	SetPageNumber(v int64) *ListChunksResponseBody
	GetPageNumber() *int64
	SetPageRecordCount(v int64) *ListChunksResponseBody
	GetPageRecordCount() *int64
	SetRequestId(v string) *ListChunksResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int64) *ListChunksResponseBody
	GetTotalRecordCount() *int64
}

type ListChunksResponseBody struct {
	Chunks *ListChunksResponseBodyChunks `json:"Chunks,omitempty" xml:"Chunks,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageRecordCount *int64 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 34b32a0a-08ef-4a87-b6be-cdd9f56fc3ad
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListChunksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChunksResponseBody) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBody) GetChunks() *ListChunksResponseBodyChunks {
	return s.Chunks
}

func (s *ListChunksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListChunksResponseBody) GetPageRecordCount() *int64 {
	return s.PageRecordCount
}

func (s *ListChunksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChunksResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *ListChunksResponseBody) SetChunks(v *ListChunksResponseBodyChunks) *ListChunksResponseBody {
	s.Chunks = v
	return s
}

func (s *ListChunksResponseBody) SetPageNumber(v int64) *ListChunksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListChunksResponseBody) SetPageRecordCount(v int64) *ListChunksResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListChunksResponseBody) SetRequestId(v string) *ListChunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChunksResponseBody) SetTotalRecordCount(v int64) *ListChunksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListChunksResponseBody) Validate() error {
	if s.Chunks != nil {
		if err := s.Chunks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChunksResponseBodyChunks struct {
	Chunks []*ListChunksResponseBodyChunksChunks `json:"chunks,omitempty" xml:"chunks,omitempty" type:"Repeated"`
}

func (s ListChunksResponseBodyChunks) String() string {
	return dara.Prettify(s)
}

func (s ListChunksResponseBodyChunks) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBodyChunks) GetChunks() []*ListChunksResponseBodyChunksChunks {
	return s.Chunks
}

func (s *ListChunksResponseBodyChunks) SetChunks(v []*ListChunksResponseBodyChunksChunks) *ListChunksResponseBodyChunks {
	s.Chunks = v
	return s
}

func (s *ListChunksResponseBodyChunks) Validate() error {
	if s.Chunks != nil {
		for _, item := range s.Chunks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChunksResponseBodyChunksChunks struct {
	Content        *string                                   `json:"Content,omitempty" xml:"Content,omitempty"`
	FileName       *string                                   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id             *string                                   `json:"Id,omitempty" xml:"Id,omitempty"`
	LoaderMetadata *string                                   `json:"LoaderMetadata,omitempty" xml:"LoaderMetadata,omitempty"`
	Metadata       map[string]interface{}                    `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	PageIndex      *int64                                    `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	Vector         *ListChunksResponseBodyChunksChunksVector `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Struct"`
}

func (s ListChunksResponseBodyChunksChunks) String() string {
	return dara.Prettify(s)
}

func (s ListChunksResponseBodyChunksChunks) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBodyChunksChunks) GetContent() *string {
	return s.Content
}

func (s *ListChunksResponseBodyChunksChunks) GetFileName() *string {
	return s.FileName
}

func (s *ListChunksResponseBodyChunksChunks) GetId() *string {
	return s.Id
}

func (s *ListChunksResponseBodyChunksChunks) GetLoaderMetadata() *string {
	return s.LoaderMetadata
}

func (s *ListChunksResponseBodyChunksChunks) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *ListChunksResponseBodyChunksChunks) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *ListChunksResponseBodyChunksChunks) GetVector() *ListChunksResponseBodyChunksChunksVector {
	return s.Vector
}

func (s *ListChunksResponseBodyChunksChunks) SetContent(v string) *ListChunksResponseBodyChunksChunks {
	s.Content = &v
	return s
}

func (s *ListChunksResponseBodyChunksChunks) SetFileName(v string) *ListChunksResponseBodyChunksChunks {
	s.FileName = &v
	return s
}

func (s *ListChunksResponseBodyChunksChunks) SetId(v string) *ListChunksResponseBodyChunksChunks {
	s.Id = &v
	return s
}

func (s *ListChunksResponseBodyChunksChunks) SetLoaderMetadata(v string) *ListChunksResponseBodyChunksChunks {
	s.LoaderMetadata = &v
	return s
}

func (s *ListChunksResponseBodyChunksChunks) SetMetadata(v map[string]interface{}) *ListChunksResponseBodyChunksChunks {
	s.Metadata = v
	return s
}

func (s *ListChunksResponseBodyChunksChunks) SetPageIndex(v int64) *ListChunksResponseBodyChunksChunks {
	s.PageIndex = &v
	return s
}

func (s *ListChunksResponseBodyChunksChunks) SetVector(v *ListChunksResponseBodyChunksChunksVector) *ListChunksResponseBodyChunksChunks {
	s.Vector = v
	return s
}

func (s *ListChunksResponseBodyChunksChunks) Validate() error {
	if s.Vector != nil {
		if err := s.Vector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChunksResponseBodyChunksChunksVector struct {
	Vector []*float64 `json:"vector,omitempty" xml:"vector,omitempty" type:"Repeated"`
}

func (s ListChunksResponseBodyChunksChunksVector) String() string {
	return dara.Prettify(s)
}

func (s ListChunksResponseBodyChunksChunksVector) GoString() string {
	return s.String()
}

func (s *ListChunksResponseBodyChunksChunksVector) GetVector() []*float64 {
	return s.Vector
}

func (s *ListChunksResponseBodyChunksChunksVector) SetVector(v []*float64) *ListChunksResponseBodyChunksChunksVector {
	s.Vector = v
	return s
}

func (s *ListChunksResponseBodyChunksChunksVector) Validate() error {
	return dara.Validate(s)
}
