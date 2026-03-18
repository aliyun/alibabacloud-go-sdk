// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPaidResourceQry interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *PaidResourceQry
	GetMaxResults() *int32
	SetNextToken(v string) *PaidResourceQry
	GetNextToken() *string
	SetResourceStatus(v string) *PaidResourceQry
	GetResourceStatus() *string
}

type PaidResourceQry struct {
	MaxResults     *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken      *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	ResourceStatus *string `json:"resourceStatus,omitempty" xml:"resourceStatus,omitempty"`
}

func (s PaidResourceQry) String() string {
	return dara.Prettify(s)
}

func (s PaidResourceQry) GoString() string {
	return s.String()
}

func (s *PaidResourceQry) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *PaidResourceQry) GetNextToken() *string {
	return s.NextToken
}

func (s *PaidResourceQry) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *PaidResourceQry) SetMaxResults(v int32) *PaidResourceQry {
	s.MaxResults = &v
	return s
}

func (s *PaidResourceQry) SetNextToken(v string) *PaidResourceQry {
	s.NextToken = &v
	return s
}

func (s *PaidResourceQry) SetResourceStatus(v string) *PaidResourceQry {
	s.ResourceStatus = &v
	return s
}

func (s *PaidResourceQry) Validate() error {
	return dara.Validate(s)
}
