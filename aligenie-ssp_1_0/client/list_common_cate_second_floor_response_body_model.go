// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCommonCateSecondFloorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCommonCateSecondFloorResponseBody
	GetCode() *int32
	SetMessage(v string) *ListCommonCateSecondFloorResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCommonCateSecondFloorResponseBody
	GetRequestId() *string
	SetResult(v []*ListCommonCateSecondFloorResponseBodyResult) *ListCommonCateSecondFloorResponseBody
	GetResult() []*ListCommonCateSecondFloorResponseBodyResult
}

type ListCommonCateSecondFloorResponseBody struct {
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
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListCommonCateSecondFloorResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListCommonCateSecondFloorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCommonCateSecondFloorResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommonCateSecondFloorResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCommonCateSecondFloorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCommonCateSecondFloorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCommonCateSecondFloorResponseBody) GetResult() []*ListCommonCateSecondFloorResponseBodyResult {
	return s.Result
}

func (s *ListCommonCateSecondFloorResponseBody) SetCode(v int32) *ListCommonCateSecondFloorResponseBody {
	s.Code = &v
	return s
}

func (s *ListCommonCateSecondFloorResponseBody) SetMessage(v string) *ListCommonCateSecondFloorResponseBody {
	s.Message = &v
	return s
}

func (s *ListCommonCateSecondFloorResponseBody) SetRequestId(v string) *ListCommonCateSecondFloorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommonCateSecondFloorResponseBody) SetResult(v []*ListCommonCateSecondFloorResponseBodyResult) *ListCommonCateSecondFloorResponseBody {
	s.Result = v
	return s
}

func (s *ListCommonCateSecondFloorResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCommonCateSecondFloorResponseBodyResult struct {
	// example:
	//
	// 80018009
	CateId   *int64  `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// example:
	//
	// 80018
	ParentCateId *int64 `json:"ParentCateId,omitempty" xml:"ParentCateId,omitempty"`
}

func (s ListCommonCateSecondFloorResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCommonCateSecondFloorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCommonCateSecondFloorResponseBodyResult) GetCateId() *int64 {
	return s.CateId
}

func (s *ListCommonCateSecondFloorResponseBodyResult) GetCateName() *string {
	return s.CateName
}

func (s *ListCommonCateSecondFloorResponseBodyResult) GetParentCateId() *int64 {
	return s.ParentCateId
}

func (s *ListCommonCateSecondFloorResponseBodyResult) SetCateId(v int64) *ListCommonCateSecondFloorResponseBodyResult {
	s.CateId = &v
	return s
}

func (s *ListCommonCateSecondFloorResponseBodyResult) SetCateName(v string) *ListCommonCateSecondFloorResponseBodyResult {
	s.CateName = &v
	return s
}

func (s *ListCommonCateSecondFloorResponseBodyResult) SetParentCateId(v int64) *ListCommonCateSecondFloorResponseBodyResult {
	s.ParentCateId = &v
	return s
}

func (s *ListCommonCateSecondFloorResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
