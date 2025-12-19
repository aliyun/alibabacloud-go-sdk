// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAPIKeyCredentialProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAPIKeyCredentialProvidersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAPIKeyCredentialProvidersRequest
	GetNextToken() *string
}

type ListAPIKeyCredentialProvidersRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l8i017XhgKLf/HqWDGD375
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAPIKeyCredentialProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAPIKeyCredentialProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListAPIKeyCredentialProvidersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAPIKeyCredentialProvidersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAPIKeyCredentialProvidersRequest) SetMaxResults(v int32) *ListAPIKeyCredentialProvidersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersRequest) SetNextToken(v string) *ListAPIKeyCredentialProvidersRequest {
	s.NextToken = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersRequest) Validate() error {
	return dara.Validate(s)
}
