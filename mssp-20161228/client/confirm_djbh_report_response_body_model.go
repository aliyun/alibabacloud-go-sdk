// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmDjbhReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConfirmDjbhReportResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *ConfirmDjbhReportResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *ConfirmDjbhReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfirmDjbhReportResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ConfirmDjbhReportResponseBody
	GetSuccess() *string
}

type ConfirmDjbhReportResponseBody struct {
	// API response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EF801DD1-D934-51B3-92D4-776CE17B184F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful. - **true**: The call was successful. - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfirmDjbhReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmDjbhReportResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmDjbhReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConfirmDjbhReportResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ConfirmDjbhReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfirmDjbhReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmDjbhReportResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ConfirmDjbhReportResponseBody) SetCode(v string) *ConfirmDjbhReportResponseBody {
	s.Code = &v
	return s
}

func (s *ConfirmDjbhReportResponseBody) SetHttpStatusCode(v string) *ConfirmDjbhReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfirmDjbhReportResponseBody) SetMessage(v string) *ConfirmDjbhReportResponseBody {
	s.Message = &v
	return s
}

func (s *ConfirmDjbhReportResponseBody) SetRequestId(v string) *ConfirmDjbhReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmDjbhReportResponseBody) SetSuccess(v string) *ConfirmDjbhReportResponseBody {
	s.Success = &v
	return s
}

func (s *ConfirmDjbhReportResponseBody) Validate() error {
	return dara.Validate(s)
}
