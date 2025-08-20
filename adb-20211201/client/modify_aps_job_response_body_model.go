// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApsJobId(v string) *ModifyApsJobResponseBody
	GetApsJobId() *string
	SetCode(v string) *ModifyApsJobResponseBody
	GetCode() *string
	SetErrCode(v string) *ModifyApsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyApsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModifyApsJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyApsJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyApsJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyApsJobResponseBody
	GetSuccess() *bool
}

type ModifyApsJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// aps-bj1xxxxxx
	ApsJobId *string `json:"ApsJobId,omitempty" xml:"ApsJobId,omitempty"`
	// The status code. A value of 200 indicates that the request is successful.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error code.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// OK
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The status code. A value of 200 indicates that the request was successful.
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
	// SUCCESS
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

func (s ModifyApsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsJobResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApsJobResponseBody) GetApsJobId() *string {
	return s.ApsJobId
}

func (s *ModifyApsJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyApsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyApsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyApsJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyApsJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyApsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApsJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyApsJobResponseBody) SetApsJobId(v string) *ModifyApsJobResponseBody {
	s.ApsJobId = &v
	return s
}

func (s *ModifyApsJobResponseBody) SetCode(v string) *ModifyApsJobResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyApsJobResponseBody) SetErrCode(v string) *ModifyApsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyApsJobResponseBody) SetErrMessage(v string) *ModifyApsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyApsJobResponseBody) SetHttpStatusCode(v int32) *ModifyApsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyApsJobResponseBody) SetMessage(v string) *ModifyApsJobResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyApsJobResponseBody) SetRequestId(v string) *ModifyApsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApsJobResponseBody) SetSuccess(v bool) *ModifyApsJobResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyApsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
