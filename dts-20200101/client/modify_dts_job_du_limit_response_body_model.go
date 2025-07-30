// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobDuLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDtsJobDuLimitResponseBody
	GetCode() *string
	SetDynamicMessage(v string) *ModifyDtsJobDuLimitResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *ModifyDtsJobDuLimitResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDtsJobDuLimitResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int64) *ModifyDtsJobDuLimitResponseBody
	GetHttpStatusCode() *int64
	SetRequestId(v string) *ModifyDtsJobDuLimitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDtsJobDuLimitResponseBody
	GetSuccess() *bool
}

type ModifyDtsJobDuLimitResponseBody struct {
	// The error code returned by the backend service. The number is incremented.
	//
	// example:
	//
	// 500
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace the %s variable in the **ErrMessage*	- parameter.
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDtsJobDuLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobDuLimitResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobDuLimitResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDtsJobDuLimitResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ModifyDtsJobDuLimitResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDtsJobDuLimitResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDtsJobDuLimitResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ModifyDtsJobDuLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDtsJobDuLimitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDtsJobDuLimitResponseBody) SetCode(v string) *ModifyDtsJobDuLimitResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDtsJobDuLimitResponseBody) SetDynamicMessage(v string) *ModifyDtsJobDuLimitResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ModifyDtsJobDuLimitResponseBody) SetErrCode(v string) *ModifyDtsJobDuLimitResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDtsJobDuLimitResponseBody) SetErrMessage(v string) *ModifyDtsJobDuLimitResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDtsJobDuLimitResponseBody) SetHttpStatusCode(v int64) *ModifyDtsJobDuLimitResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDtsJobDuLimitResponseBody) SetRequestId(v string) *ModifyDtsJobDuLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDtsJobDuLimitResponseBody) SetSuccess(v bool) *ModifyDtsJobDuLimitResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDtsJobDuLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
