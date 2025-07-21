// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportUserFbIssueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReportUserFbIssueResponseBody
	GetCode() *string
	SetData(v *ReportUserFbIssueResponseBodyData) *ReportUserFbIssueResponseBody
	GetData() *ReportUserFbIssueResponseBodyData
	SetMessage(v string) *ReportUserFbIssueResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReportUserFbIssueResponseBody
	GetRequestId() *string
}

type ReportUserFbIssueResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ReportUserFbIssueResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportUserFbIssueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbIssueResponseBody) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReportUserFbIssueResponseBody) GetData() *ReportUserFbIssueResponseBodyData {
	return s.Data
}

func (s *ReportUserFbIssueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReportUserFbIssueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportUserFbIssueResponseBody) SetCode(v string) *ReportUserFbIssueResponseBody {
	s.Code = &v
	return s
}

func (s *ReportUserFbIssueResponseBody) SetData(v *ReportUserFbIssueResponseBodyData) *ReportUserFbIssueResponseBody {
	s.Data = v
	return s
}

func (s *ReportUserFbIssueResponseBody) SetMessage(v string) *ReportUserFbIssueResponseBody {
	s.Message = &v
	return s
}

func (s *ReportUserFbIssueResponseBody) SetRequestId(v string) *ReportUserFbIssueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportUserFbIssueResponseBody) Validate() error {
	return dara.Validate(s)
}

type ReportUserFbIssueResponseBodyData struct {
	IssueId *int32 `json:"IssueId,omitempty" xml:"IssueId,omitempty"`
}

func (s ReportUserFbIssueResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbIssueResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueResponseBodyData) GetIssueId() *int32 {
	return s.IssueId
}

func (s *ReportUserFbIssueResponseBodyData) SetIssueId(v int32) *ReportUserFbIssueResponseBodyData {
	s.IssueId = &v
	return s
}

func (s *ReportUserFbIssueResponseBodyData) Validate() error {
	return dara.Validate(s)
}
