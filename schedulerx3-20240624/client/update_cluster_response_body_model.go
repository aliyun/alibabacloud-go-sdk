// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateClusterResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateClusterResponseBody
	GetSuccess() *bool
}

type UpdateClusterResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEF9AF15-AEEF-5E59-BF7B-BCBB119DC53F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateClusterResponseBody) SetCode(v int32) *UpdateClusterResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateClusterResponseBody) SetMessage(v string) *UpdateClusterResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateClusterResponseBody) SetRequestId(v string) *UpdateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterResponseBody) SetSuccess(v bool) *UpdateClusterResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
