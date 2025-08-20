// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitResultExportJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitResultExportJobResponseBody
	GetCode() *string
	SetData(v string) *SubmitResultExportJobResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *SubmitResultExportJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitResultExportJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitResultExportJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitResultExportJobResponseBody
	GetSuccess() *bool
}

type SubmitResultExportJobResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the result set export job.
	//
	// example:
	//
	// export_202405131927121980210080040****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, an **OK*	- message is returned.
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
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitResultExportJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitResultExportJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitResultExportJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitResultExportJobResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitResultExportJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitResultExportJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitResultExportJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitResultExportJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitResultExportJobResponseBody) SetCode(v string) *SubmitResultExportJobResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitResultExportJobResponseBody) SetData(v string) *SubmitResultExportJobResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitResultExportJobResponseBody) SetHttpStatusCode(v int32) *SubmitResultExportJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitResultExportJobResponseBody) SetMessage(v string) *SubmitResultExportJobResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitResultExportJobResponseBody) SetRequestId(v string) *SubmitResultExportJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitResultExportJobResponseBody) SetSuccess(v bool) *SubmitResultExportJobResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitResultExportJobResponseBody) Validate() error {
	return dara.Validate(s)
}
