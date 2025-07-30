// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDtsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *ModifyDtsJobResponseBody
	GetDtsJobId() *string
	SetErrCode(v string) *ModifyDtsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v bool) *ModifyDtsJobResponseBody
	GetErrMessage() *bool
	SetRequestId(v string) *ModifyDtsJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *ModifyDtsJobResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ModifyDtsJobResponseBody
	GetSuccess() *bool
}

type ModifyDtsJobResponseBody struct {
	// The ID of the DTS task.
	//
	// example:
	//
	// bi6e22a****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// >  The data type of this parameter is String. Sample value: **The actual sample value is The request processing has failed due to some unknown error.
	//
	// example:
	//
	// true
	ErrMessage *bool `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1D6ECADF-C5E9-4C96-8811-77602B31****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDtsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifyDtsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDtsJobResponseBody) GetErrMessage() *bool {
	return s.ErrMessage
}

func (s *ModifyDtsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDtsJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ModifyDtsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDtsJobResponseBody) SetDtsJobId(v string) *ModifyDtsJobResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobResponseBody) SetErrCode(v string) *ModifyDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDtsJobResponseBody) SetErrMessage(v bool) *ModifyDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDtsJobResponseBody) SetRequestId(v string) *ModifyDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDtsJobResponseBody) SetStatus(v string) *ModifyDtsJobResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyDtsJobResponseBody) SetSuccess(v bool) *ModifyDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDtsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
