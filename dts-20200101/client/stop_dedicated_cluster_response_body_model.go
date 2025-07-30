// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDedicatedClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *StopDedicatedClusterResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StopDedicatedClusterResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *StopDedicatedClusterResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *StopDedicatedClusterResponseBody
	GetRequestId() *string
	SetSuccess(v string) *StopDedicatedClusterResponseBody
	GetSuccess() *string
}

type StopDedicatedClusterResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.q
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopDedicatedClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDedicatedClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopDedicatedClusterResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StopDedicatedClusterResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StopDedicatedClusterResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *StopDedicatedClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDedicatedClusterResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *StopDedicatedClusterResponseBody) SetErrCode(v string) *StopDedicatedClusterResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StopDedicatedClusterResponseBody) SetErrMessage(v string) *StopDedicatedClusterResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StopDedicatedClusterResponseBody) SetHttpStatusCode(v string) *StopDedicatedClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopDedicatedClusterResponseBody) SetRequestId(v string) *StopDedicatedClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDedicatedClusterResponseBody) SetSuccess(v string) *StopDedicatedClusterResponseBody {
	s.Success = &v
	return s
}

func (s *StopDedicatedClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
