// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReverseDtsJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *CreateReverseDtsJobResponseBody
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *CreateReverseDtsJobResponseBody
	GetDtsJobId() *string
	SetErrCode(v string) *CreateReverseDtsJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateReverseDtsJobResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *CreateReverseDtsJobResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateReverseDtsJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateReverseDtsJobResponseBody
	GetSuccess() *string
}

type CreateReverseDtsJobResponseBody struct {
	// The ID of the instance corresponding to the generated reverse task.
	//
	// example:
	//
	// dtsor1f9kr822l****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the generated reverse synchronization task.
	//
	// example:
	//
	// n99m9jx822k****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// Error code returned when the call fails.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// Error message returned when the call fails.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// Status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateReverseDtsJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReverseDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReverseDtsJobResponseBody) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *CreateReverseDtsJobResponseBody) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *CreateReverseDtsJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateReverseDtsJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateReverseDtsJobResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateReverseDtsJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReverseDtsJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateReverseDtsJobResponseBody) SetDtsInstanceId(v string) *CreateReverseDtsJobResponseBody {
	s.DtsInstanceId = &v
	return s
}

func (s *CreateReverseDtsJobResponseBody) SetDtsJobId(v string) *CreateReverseDtsJobResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *CreateReverseDtsJobResponseBody) SetErrCode(v string) *CreateReverseDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateReverseDtsJobResponseBody) SetErrMessage(v string) *CreateReverseDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateReverseDtsJobResponseBody) SetHttpStatusCode(v string) *CreateReverseDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateReverseDtsJobResponseBody) SetRequestId(v string) *CreateReverseDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReverseDtsJobResponseBody) SetSuccess(v string) *CreateReverseDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateReverseDtsJobResponseBody) Validate() error {
	return dara.Validate(s)
}
