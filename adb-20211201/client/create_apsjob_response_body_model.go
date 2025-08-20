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
	// The job ID.
	//
	// example:
	//
	// aps-bj1xxxxxx
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The HTTP status code or the error code.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D1B8ED33-5E9B-512D-B188-1579ED6xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
