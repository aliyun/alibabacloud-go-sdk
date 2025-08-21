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
	// The number of entries per page. If a response is truncated because it reaches the value of `MaxItems`, the value of `IsTruncated` will be true.
	//
	// Valid values: 1 to 3501. Default value: 3501.
	//
	// example:
	//
	// 1000
	MaxItems *string `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
	// The token that is used to initiate the next request if the response of the current request is truncated. You can use the token to initiate another request and obtain the remaining records.``
	//
	// example:
	//
	// EXAMPLE
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
