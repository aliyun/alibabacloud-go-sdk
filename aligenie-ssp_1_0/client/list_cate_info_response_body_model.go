// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCateInfoResponseBody
	GetCode() *int32
	SetMessage(v string) *ListCateInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCateInfoResponseBody
	GetRequestId() *string
	SetResult(v []*ListCateInfoResponseBodyResult) *ListCateInfoResponseBody
	GetResult() []*ListCateInfoResponseBodyResult
}

type ListCateInfoResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F12B6147-5925-19E5-A3AD-E1EE1360F34E
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListCateInfoResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListCateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListCateInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCateInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCateInfoResponseBody) GetResult() []*ListCateInfoResponseBodyResult {
	return s.Result
}

func (s *ListCateInfoResponseBody) SetCode(v int32) *ListCateInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListCateInfoResponseBody) SetMessage(v string) *ListCateInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListCateInfoResponseBody) SetRequestId(v string) *ListCateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCateInfoResponseBody) SetResult(v []*ListCateInfoResponseBodyResult) *ListCateInfoResponseBody {
	s.Result = v
	return s
}

func (s *ListCateInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCateInfoResponseBodyResult struct {
	// example:
	//
	// 80064
	CateId   *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// example:
	//
	// 0
	ParentCateId *int64 `json:"ParentCateId,omitempty" xml:"ParentCateId,omitempty"`
}

func (s ListCateInfoResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCateInfoResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCateInfoResponseBodyResult) GetCateId() *int64 {
	return s.CateId
}

func (s *ListCateInfoResponseBodyResult) GetCateName() *string {
	return s.CateName
}

func (s *ListCateInfoResponseBodyResult) GetParentCateId() *int64 {
	return s.ParentCateId
}

func (s *ListCateInfoResponseBodyResult) SetCateId(v int64) *ListCateInfoResponseBodyResult {
	s.CateId = &v
	return s
}

func (s *ListCateInfoResponseBodyResult) SetCateName(v string) *ListCateInfoResponseBodyResult {
	s.CateName = &v
	return s
}

func (s *ListCateInfoResponseBodyResult) SetParentCateId(v int64) *ListCateInfoResponseBodyResult {
	s.ParentCateId = &v
	return s
}

func (s *ListCateInfoResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
