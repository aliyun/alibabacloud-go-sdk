// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonCateFirstFloorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCommonCateFirstFloorResponseBody
	GetCode() *int32
	SetMessage(v string) *ListCommonCateFirstFloorResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCommonCateFirstFloorResponseBody
	GetRequestId() *string
	SetResult(v []*ListCommonCateFirstFloorResponseBodyResult) *ListCommonCateFirstFloorResponseBody
	GetResult() []*ListCommonCateFirstFloorResponseBodyResult
}

type ListCommonCateFirstFloorResponseBody struct {
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
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListCommonCateFirstFloorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListCommonCateFirstFloorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCommonCateFirstFloorResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommonCateFirstFloorResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCommonCateFirstFloorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCommonCateFirstFloorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCommonCateFirstFloorResponseBody) GetResult() []*ListCommonCateFirstFloorResponseBodyResult {
	return s.Result
}

func (s *ListCommonCateFirstFloorResponseBody) SetCode(v int32) *ListCommonCateFirstFloorResponseBody {
	s.Code = &v
	return s
}

func (s *ListCommonCateFirstFloorResponseBody) SetMessage(v string) *ListCommonCateFirstFloorResponseBody {
	s.Message = &v
	return s
}

func (s *ListCommonCateFirstFloorResponseBody) SetRequestId(v string) *ListCommonCateFirstFloorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommonCateFirstFloorResponseBody) SetResult(v []*ListCommonCateFirstFloorResponseBodyResult) *ListCommonCateFirstFloorResponseBody {
	s.Result = v
	return s
}

func (s *ListCommonCateFirstFloorResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCommonCateFirstFloorResponseBodyResult struct {
	// example:
	//
	// 80012
	CateId   *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// example:
	//
	// 0
	ParentCateId *int64 `json:"ParentCateId,omitempty" xml:"ParentCateId,omitempty"`
}

func (s ListCommonCateFirstFloorResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCommonCateFirstFloorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCommonCateFirstFloorResponseBodyResult) GetCateId() *int64 {
	return s.CateId
}

func (s *ListCommonCateFirstFloorResponseBodyResult) GetCateName() *string {
	return s.CateName
}

func (s *ListCommonCateFirstFloorResponseBodyResult) GetParentCateId() *int64 {
	return s.ParentCateId
}

func (s *ListCommonCateFirstFloorResponseBodyResult) SetCateId(v int64) *ListCommonCateFirstFloorResponseBodyResult {
	s.CateId = &v
	return s
}

func (s *ListCommonCateFirstFloorResponseBodyResult) SetCateName(v string) *ListCommonCateFirstFloorResponseBodyResult {
	s.CateName = &v
	return s
}

func (s *ListCommonCateFirstFloorResponseBodyResult) SetParentCateId(v int64) *ListCommonCateFirstFloorResponseBodyResult {
	s.ParentCateId = &v
	return s
}

func (s *ListCommonCateFirstFloorResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
