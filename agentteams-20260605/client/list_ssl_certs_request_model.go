// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSslCertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListSslCertsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListSslCertsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSslCertsRequest
	GetNextToken() *string
}

type ListSslCertsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// instance-1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListSslCertsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSslCertsRequest) GoString() string {
	return s.String()
}

func (s *ListSslCertsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSslCertsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSslCertsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSslCertsRequest) SetInstanceId(v string) *ListSslCertsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSslCertsRequest) SetMaxResults(v int32) *ListSslCertsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSslCertsRequest) SetNextToken(v string) *ListSslCertsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSslCertsRequest) Validate() error {
	return dara.Validate(s)
}
