// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSceneCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSceneCategoryResponseBody
	GetCode() *int32
	SetMessage(v string) *ListSceneCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSceneCategoryResponseBody
	GetRequestId() *string
	SetResult(v []*string) *ListSceneCategoryResponseBody
	GetResult() []*string
}

type ListSceneCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	//
	// example:
	//
	// 0EC7*726E
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListSceneCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSceneCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListSceneCategoryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSceneCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSceneCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSceneCategoryResponseBody) GetResult() []*string {
	return s.Result
}

func (s *ListSceneCategoryResponseBody) SetCode(v int32) *ListSceneCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListSceneCategoryResponseBody) SetMessage(v string) *ListSceneCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListSceneCategoryResponseBody) SetRequestId(v string) *ListSceneCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSceneCategoryResponseBody) SetResult(v []*string) *ListSceneCategoryResponseBody {
	s.Result = v
	return s
}

func (s *ListSceneCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
