// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportUserFbAcIssueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReportUserFbAcIssueResponseBody
	GetCode() *string
	SetData(v *ReportUserFbAcIssueResponseBodyData) *ReportUserFbAcIssueResponseBody
	GetData() *ReportUserFbAcIssueResponseBodyData
	SetMessage(v string) *ReportUserFbAcIssueResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReportUserFbAcIssueResponseBody
	GetRequestId() *string
}

type ReportUserFbAcIssueResponseBody struct {
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ReportUserFbAcIssueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportUserFbAcIssueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbAcIssueResponseBody) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReportUserFbAcIssueResponseBody) GetData() *ReportUserFbAcIssueResponseBodyData {
	return s.Data
}

func (s *ReportUserFbAcIssueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReportUserFbAcIssueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportUserFbAcIssueResponseBody) SetCode(v string) *ReportUserFbAcIssueResponseBody {
	s.Code = &v
	return s
}

func (s *ReportUserFbAcIssueResponseBody) SetData(v *ReportUserFbAcIssueResponseBodyData) *ReportUserFbAcIssueResponseBody {
	s.Data = v
	return s
}

func (s *ReportUserFbAcIssueResponseBody) SetMessage(v string) *ReportUserFbAcIssueResponseBody {
	s.Message = &v
	return s
}

func (s *ReportUserFbAcIssueResponseBody) SetRequestId(v string) *ReportUserFbAcIssueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportUserFbAcIssueResponseBody) Validate() error {
	return dara.Validate(s)
}

type ReportUserFbAcIssueResponseBodyData struct {
	IssueId *int64 `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
}

func (s ReportUserFbAcIssueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbAcIssueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueResponseBodyData) GetIssueId() *int64 {
	return s.IssueId
}

func (s *ReportUserFbAcIssueResponseBodyData) SetIssueId(v int64) *ReportUserFbAcIssueResponseBodyData {
	s.IssueId = &v
	return s
}

func (s *ReportUserFbAcIssueResponseBodyData) Validate() error {
	return dara.Validate(s)
}
