// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetCredentialReportResponseBody
	GetContent() *string
	SetGeneratedTime(v string) *GetCredentialReportResponseBody
	GetGeneratedTime() *string
	SetIsTruncated(v string) *GetCredentialReportResponseBody
	GetIsTruncated() *string
	SetNextToken(v string) *GetCredentialReportResponseBody
	GetNextToken() *string
	SetRequestId(v string) *GetCredentialReportResponseBody
	GetRequestId() *string
}

type GetCredentialReportResponseBody struct {
	// The content of the user credential report.
	//
	// The report is Base64-encoded. After you decode the report, the credential report is in the CSV format.
	//
	// example:
	//
	// OVZWK4RMOVZW****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time when the user credential report was generated.
	//
	// example:
	//
	// 2020-10-19T15:06:52Z
	GeneratedTime *string `json:"GeneratedTime,omitempty" xml:"GeneratedTime,omitempty"`
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *string `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The parameter that is used to obtain the truncated part. This parameter takes effect only when `IsTruncated` is set to true.
	//
	// example:
	//
	// EXAMPLE
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7A01826E-7601-44B0-B4DF-2B0C509836DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCredentialReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetCredentialReportResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetCredentialReportResponseBody) GetGeneratedTime() *string {
	return s.GeneratedTime
}

func (s *GetCredentialReportResponseBody) GetIsTruncated() *string {
	return s.IsTruncated
}

func (s *GetCredentialReportResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetCredentialReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCredentialReportResponseBody) SetContent(v string) *GetCredentialReportResponseBody {
	s.Content = &v
	return s
}

func (s *GetCredentialReportResponseBody) SetGeneratedTime(v string) *GetCredentialReportResponseBody {
	s.GeneratedTime = &v
	return s
}

func (s *GetCredentialReportResponseBody) SetIsTruncated(v string) *GetCredentialReportResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *GetCredentialReportResponseBody) SetNextToken(v string) *GetCredentialReportResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetCredentialReportResponseBody) SetRequestId(v string) *GetCredentialReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCredentialReportResponseBody) Validate() error {
	return dara.Validate(s)
}
