// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobDedicatedClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDtsJobDedicatedClusterResponseBody
	GetCode() *string
	SetDynamicMessage(v string) *ModifyDtsJobDedicatedClusterResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *ModifyDtsJobDedicatedClusterResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDtsJobDedicatedClusterResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int64) *ModifyDtsJobDedicatedClusterResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *ModifyDtsJobDedicatedClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDtsJobDedicatedClusterResponseBody
	GetSuccess() *bool
}

type ModifyDtsJobDedicatedClusterResponseBody struct {
	// The error code returned by the backend service.
	//
	// example:
	//
	// 500
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace %s in **ErrMessage**.
	//
	// example:
	//
	// Type
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
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
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDtsJobDedicatedClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobDedicatedClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) SetCode(v string) *ModifyDtsJobDedicatedClusterResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) SetDynamicMessage(v string) *ModifyDtsJobDedicatedClusterResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) SetErrCode(v string) *ModifyDtsJobDedicatedClusterResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) SetErrMessage(v string) *ModifyDtsJobDedicatedClusterResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) SetHttpStatusCode(v int64) *ModifyDtsJobDedicatedClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) SetRequestId(v string) *ModifyDtsJobDedicatedClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) SetSuccess(v bool) *ModifyDtsJobDedicatedClusterResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDtsJobDedicatedClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
