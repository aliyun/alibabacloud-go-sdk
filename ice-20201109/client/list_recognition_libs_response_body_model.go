// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecognitionLibsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLibs(v *ListRecognitionLibsResponseBodyLibs) *ListRecognitionLibsResponseBody
	GetLibs() *ListRecognitionLibsResponseBodyLibs
	SetPageNumber(v int32) *ListRecognitionLibsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecognitionLibsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRecognitionLibsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRecognitionLibsResponseBody
	GetTotalCount() *int64
}

type ListRecognitionLibsResponseBody struct {
	// The recognition libraries.
	Libs *ListRecognitionLibsResponseBodyLibs `json:"Libs,omitempty" xml:"Libs,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 180
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecognitionLibsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionLibsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecognitionLibsResponseBody) GetLibs() *ListRecognitionLibsResponseBodyLibs {
	return s.Libs
}

func (s *ListRecognitionLibsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecognitionLibsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecognitionLibsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecognitionLibsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRecognitionLibsResponseBody) SetLibs(v *ListRecognitionLibsResponseBodyLibs) *ListRecognitionLibsResponseBody {
	s.Libs = v
	return s
}

func (s *ListRecognitionLibsResponseBody) SetPageNumber(v int32) *ListRecognitionLibsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRecognitionLibsResponseBody) SetPageSize(v int32) *ListRecognitionLibsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRecognitionLibsResponseBody) SetRequestId(v string) *ListRecognitionLibsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecognitionLibsResponseBody) SetTotalCount(v int64) *ListRecognitionLibsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecognitionLibsResponseBody) Validate() error {
	if s.Libs != nil {
		if err := s.Libs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRecognitionLibsResponseBodyLibs struct {
	Lib []*ListRecognitionLibsResponseBodyLibsLib `json:"Lib,omitempty" xml:"Lib,omitempty" type:"Repeated"`
}

func (s ListRecognitionLibsResponseBodyLibs) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionLibsResponseBodyLibs) GoString() string {
	return s.String()
}

func (s *ListRecognitionLibsResponseBodyLibs) GetLib() []*ListRecognitionLibsResponseBodyLibsLib {
	return s.Lib
}

func (s *ListRecognitionLibsResponseBodyLibs) SetLib(v []*ListRecognitionLibsResponseBodyLibsLib) *ListRecognitionLibsResponseBodyLibs {
	s.Lib = v
	return s
}

func (s *ListRecognitionLibsResponseBodyLibs) Validate() error {
	if s.Lib != nil {
		for _, item := range s.Lib {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecognitionLibsResponseBodyLibsLib struct {
	// The description of the recognition library.
	LibDescription *string `json:"LibDescription,omitempty" xml:"LibDescription,omitempty"`
	// The ID of the recognition library.
	//
	// example:
	//
	// *************24b47865c6**************
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The name of the recognition library.
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ListRecognitionLibsResponseBodyLibsLib) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionLibsResponseBodyLibsLib) GoString() string {
	return s.String()
}

func (s *ListRecognitionLibsResponseBodyLibsLib) GetLibDescription() *string {
	return s.LibDescription
}

func (s *ListRecognitionLibsResponseBodyLibsLib) GetLibId() *string {
	return s.LibId
}

func (s *ListRecognitionLibsResponseBodyLibsLib) GetLibName() *string {
	return s.LibName
}

func (s *ListRecognitionLibsResponseBodyLibsLib) SetLibDescription(v string) *ListRecognitionLibsResponseBodyLibsLib {
	s.LibDescription = &v
	return s
}

func (s *ListRecognitionLibsResponseBodyLibsLib) SetLibId(v string) *ListRecognitionLibsResponseBodyLibsLib {
	s.LibId = &v
	return s
}

func (s *ListRecognitionLibsResponseBodyLibsLib) SetLibName(v string) *ListRecognitionLibsResponseBodyLibsLib {
	s.LibName = &v
	return s
}

func (s *ListRecognitionLibsResponseBodyLibsLib) Validate() error {
	return dara.Validate(s)
}
