// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReverseTwoWayDirectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReverseTwoWayDirectionResponseBody
	GetCode() *string
	SetDynamicMessage(v string) *ReverseTwoWayDirectionResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *ReverseTwoWayDirectionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ReverseTwoWayDirectionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ReverseTwoWayDirectionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ReverseTwoWayDirectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReverseTwoWayDirectionResponseBody
	GetSuccess() *bool
}

type ReverseTwoWayDirectionResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// DtsInstanceId
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// example:
	//
	// 403
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// AD823BD3-1BA6-4117-A536-165CB280****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReverseTwoWayDirectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReverseTwoWayDirectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReverseTwoWayDirectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReverseTwoWayDirectionResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ReverseTwoWayDirectionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ReverseTwoWayDirectionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ReverseTwoWayDirectionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ReverseTwoWayDirectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReverseTwoWayDirectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReverseTwoWayDirectionResponseBody) SetCode(v string) *ReverseTwoWayDirectionResponseBody {
	s.Code = &v
	return s
}

func (s *ReverseTwoWayDirectionResponseBody) SetDynamicMessage(v string) *ReverseTwoWayDirectionResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ReverseTwoWayDirectionResponseBody) SetErrCode(v string) *ReverseTwoWayDirectionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ReverseTwoWayDirectionResponseBody) SetErrMessage(v string) *ReverseTwoWayDirectionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ReverseTwoWayDirectionResponseBody) SetHttpStatusCode(v int32) *ReverseTwoWayDirectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ReverseTwoWayDirectionResponseBody) SetRequestId(v string) *ReverseTwoWayDirectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReverseTwoWayDirectionResponseBody) SetSuccess(v bool) *ReverseTwoWayDirectionResponseBody {
	s.Success = &v
	return s
}

func (s *ReverseTwoWayDirectionResponseBody) Validate() error {
	return dara.Validate(s)
}
