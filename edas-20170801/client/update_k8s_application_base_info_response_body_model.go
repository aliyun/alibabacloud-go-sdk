// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sApplicationBaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateK8sApplicationBaseInfoResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateK8sApplicationBaseInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateK8sApplicationBaseInfoResponseBody
	GetRequestId() *string
	SetResult(v string) *UpdateK8sApplicationBaseInfoResponseBody
	GetResult() *string
}

type UpdateK8sApplicationBaseInfoResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 57F146F6-3C94-****-****-A66EF4B9*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the modification.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateK8sApplicationBaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sApplicationBaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) GetResult() *string {
	return s.Result
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) SetCode(v int32) *UpdateK8sApplicationBaseInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) SetMessage(v string) *UpdateK8sApplicationBaseInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) SetRequestId(v string) *UpdateK8sApplicationBaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) SetResult(v string) *UpdateK8sApplicationBaseInfoResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateK8sApplicationBaseInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
