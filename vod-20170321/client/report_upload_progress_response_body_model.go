// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportUploadProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ReportUploadProgressResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReportUploadProgressResponseBody
	GetRequestId() *string
}

type ReportUploadProgressResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportUploadProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportUploadProgressResponseBody) GoString() string {
	return s.String()
}

func (s *ReportUploadProgressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReportUploadProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportUploadProgressResponseBody) SetMessage(v string) *ReportUploadProgressResponseBody {
	s.Message = &v
	return s
}

func (s *ReportUploadProgressResponseBody) SetRequestId(v string) *ReportUploadProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportUploadProgressResponseBody) Validate() error {
	return dara.Validate(s)
}
