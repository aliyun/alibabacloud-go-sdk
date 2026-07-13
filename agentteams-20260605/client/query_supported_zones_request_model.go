// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupportedZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *QuerySupportedZonesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *QuerySupportedZonesRequest
	GetNextToken() *string
}

type QuerySupportedZonesRequest struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s QuerySupportedZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySupportedZonesRequest) GoString() string {
	return s.String()
}

func (s *QuerySupportedZonesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QuerySupportedZonesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QuerySupportedZonesRequest) SetMaxResults(v int32) *QuerySupportedZonesRequest {
	s.MaxResults = &v
	return s
}

func (s *QuerySupportedZonesRequest) SetNextToken(v string) *QuerySupportedZonesRequest {
	s.NextToken = &v
	return s
}

func (s *QuerySupportedZonesRequest) Validate() error {
	return dara.Validate(s)
}
