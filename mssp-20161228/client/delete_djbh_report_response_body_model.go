// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDjbhReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDjbhReportResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DeleteDjbhReportResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DeleteDjbhReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDjbhReportResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteDjbhReportResponseBody
	GetSuccess() *string
}

type DeleteDjbhReportResponseBody struct {
	// API response code.
	//
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 86786E4C-6416-55CF-9AB6-5E275B68801D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDjbhReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDjbhReportResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDjbhReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDjbhReportResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DeleteDjbhReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDjbhReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDjbhReportResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteDjbhReportResponseBody) SetCode(v string) *DeleteDjbhReportResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDjbhReportResponseBody) SetHttpStatusCode(v string) *DeleteDjbhReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDjbhReportResponseBody) SetMessage(v string) *DeleteDjbhReportResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDjbhReportResponseBody) SetRequestId(v string) *DeleteDjbhReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDjbhReportResponseBody) SetSuccess(v string) *DeleteDjbhReportResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDjbhReportResponseBody) Validate() error {
	return dara.Validate(s)
}
