// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeGroupDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *UpdateNodeGroupDescriptionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateNodeGroupDescriptionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdateNodeGroupDescriptionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateNodeGroupDescriptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNodeGroupDescriptionResponseBody
	GetSuccess() *bool
}

type UpdateNodeGroupDescriptionResponseBody struct {
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// null
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNodeGroupDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupDescriptionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateNodeGroupDescriptionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateNodeGroupDescriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateNodeGroupDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNodeGroupDescriptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNodeGroupDescriptionResponseBody) SetErrCode(v string) *UpdateNodeGroupDescriptionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateNodeGroupDescriptionResponseBody) SetErrMessage(v string) *UpdateNodeGroupDescriptionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateNodeGroupDescriptionResponseBody) SetHttpStatusCode(v int32) *UpdateNodeGroupDescriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateNodeGroupDescriptionResponseBody) SetRequestId(v string) *UpdateNodeGroupDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeGroupDescriptionResponseBody) SetSuccess(v bool) *UpdateNodeGroupDescriptionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNodeGroupDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
