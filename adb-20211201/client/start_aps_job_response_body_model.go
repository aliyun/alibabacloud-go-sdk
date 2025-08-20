// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartApsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobId(v string) *StartApsJobResponseBody
	GetApsJobId() *string
	SetCode(v string) *StartApsJobResponseBody
	GetCode() *string
	SetErrCode(v string) *StartApsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StartApsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *StartApsJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartApsJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartApsJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartApsJobResponseBody
	GetSuccess() *bool
}

type StartApsJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// aps-******
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The HTTP status code or the error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// 0
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// OK
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
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
	// ******-3EEC-******-9F06-******
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

func (s StartApsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartApsJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartApsJobResponseBody) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *StartApsJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartApsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StartApsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StartApsJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartApsJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartApsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartApsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartApsJobResponseBody) SetApsJobId(v string) *StartApsJobResponseBody {
	s.ApsJobId = &v
	return s
}

func (s *StartApsJobResponseBody) SetCode(v string) *StartApsJobResponseBody {
	s.Code = &v
	return s
}

func (s *StartApsJobResponseBody) SetErrCode(v string) *StartApsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartApsJobResponseBody) SetErrMessage(v string) *StartApsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartApsJobResponseBody) SetHttpStatusCode(v int32) *StartApsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartApsJobResponseBody) SetMessage(v string) *StartApsJobResponseBody {
	s.Message = &v
	return s
}

func (s *StartApsJobResponseBody) SetRequestId(v string) *StartApsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartApsJobResponseBody) SetSuccess(v bool) *StartApsJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartApsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
