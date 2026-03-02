// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxItems(v string) *GetCredentialReportRequest
	GetMaxItems() *string
	SetNextToken(v string) *GetCredentialReportRequest
	GetNextToken() *string
}

type GetCredentialReportRequest struct {
	MaxItems  *string `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetCredentialReportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialReportRequest) GoString() string {
	return s.String()
}

func (s *GetCredentialReportRequest) GetMaxItems() *string {
	return s.MaxItems
}

func (s *GetCredentialReportRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetCredentialReportRequest) SetMaxItems(v string) *GetCredentialReportRequest {
	s.MaxItems = &v
	return s
}

func (s *GetCredentialReportRequest) SetNextToken(v string) *GetCredentialReportRequest {
	s.NextToken = &v
	return s
}

func (s *GetCredentialReportRequest) Validate() error {
	return dara.Validate(s)
}
