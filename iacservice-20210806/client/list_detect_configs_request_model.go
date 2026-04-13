// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDetectConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetectConfigName(v string) *ListDetectConfigsRequest
	GetDetectConfigName() *string
	SetMaxResults(v int32) *ListDetectConfigsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDetectConfigsRequest
	GetNextToken() *string
}

type ListDetectConfigsRequest struct {
	// example:
	//
	// test
	DetectConfigName *string `json:"detectConfigName,omitempty" xml:"detectConfigName,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 30BaZ9ekYWXJdqshYecA++coNg7qT1Zbm3RfLyFIZeY=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListDetectConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDetectConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListDetectConfigsRequest) GetDetectConfigName() *string {
	return s.DetectConfigName
}

func (s *ListDetectConfigsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDetectConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDetectConfigsRequest) SetDetectConfigName(v string) *ListDetectConfigsRequest {
	s.DetectConfigName = &v
	return s
}

func (s *ListDetectConfigsRequest) SetMaxResults(v int32) *ListDetectConfigsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDetectConfigsRequest) SetNextToken(v string) *ListDetectConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDetectConfigsRequest) Validate() error {
	return dara.Validate(s)
}
