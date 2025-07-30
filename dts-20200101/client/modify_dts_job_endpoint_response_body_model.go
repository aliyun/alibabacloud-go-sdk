// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ModifyDtsJobEndpointResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDtsJobEndpointResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *ModifyDtsJobEndpointResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *ModifyDtsJobEndpointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDtsJobEndpointResponseBody
	GetSuccess() *bool
}

type ModifyDtsJobEndpointResponseBody struct {
	// Error code returned when the call fails.
	//
	// example:
	//
	// DTS.Msg.InvalidEndpoint
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// Error message returned when the call fails.
	//
	// example:
	//
	// The endpoint is invalid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// request ID
	//
	// example:
	//
	// 3FA98DF2-2F81-51FF-8A38-AA5112DD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDtsJobEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobEndpointResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDtsJobEndpointResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDtsJobEndpointResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ModifyDtsJobEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDtsJobEndpointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDtsJobEndpointResponseBody) SetErrCode(v string) *ModifyDtsJobEndpointResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDtsJobEndpointResponseBody) SetErrMessage(v string) *ModifyDtsJobEndpointResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDtsJobEndpointResponseBody) SetHttpStatusCode(v string) *ModifyDtsJobEndpointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDtsJobEndpointResponseBody) SetRequestId(v string) *ModifyDtsJobEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDtsJobEndpointResponseBody) SetSuccess(v bool) *ModifyDtsJobEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDtsJobEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}
