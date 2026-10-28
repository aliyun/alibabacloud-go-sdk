// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAPSJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobId(v string) *CreateAPSJobResponseBody
	GetApsJobId() *string
	SetCode(v string) *CreateAPSJobResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateAPSJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAPSJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAPSJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAPSJobResponseBody
	GetSuccess() *bool
}

type CreateAPSJobResponseBody struct {
	// The ID of the task.
	//
	// example:
	//
	// aps-bj1xxxxxx
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The operation status code or POP error code.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned. A value of 200 indicates that the request was sent successfully.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information about the call result. Valid values:
	//
	// - If the request is normal, **OK*	- is returned.
	//
	// - If the request is abnormal, a specific error code is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D1B8ED33-5E9B-512D-B188-1579ED6xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAPSJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAPSJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAPSJobResponseBody) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *CreateAPSJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAPSJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAPSJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAPSJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAPSJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAPSJobResponseBody) SetApsJobId(v string) *CreateAPSJobResponseBody {
	s.ApsJobId = &v
	return s
}

func (s *CreateAPSJobResponseBody) SetCode(v string) *CreateAPSJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAPSJobResponseBody) SetHttpStatusCode(v int32) *CreateAPSJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAPSJobResponseBody) SetMessage(v string) *CreateAPSJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAPSJobResponseBody) SetRequestId(v string) *CreateAPSJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAPSJobResponseBody) SetSuccess(v bool) *CreateAPSJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAPSJobResponseBody) Validate() error {
	return dara.Validate(s)
}
