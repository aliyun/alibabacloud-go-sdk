// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFederationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFederations(v []*ListFederationsResponseBodyFederations) *ListFederationsResponseBody
	GetFederations() []*ListFederationsResponseBodyFederations
	SetMaxResults(v int32) *ListFederationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListFederationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListFederationsResponseBody
	GetRequestId() *string
}

type ListFederationsResponseBody struct {
	Federations []*ListFederationsResponseBodyFederations `json:"Federations,omitempty" xml:"Federations,omitempty" type:"Repeated"`
	MaxResults  *int32                                    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken   *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFederationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFederationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFederationsResponseBody) GetFederations() []*ListFederationsResponseBodyFederations {
	return s.Federations
}

func (s *ListFederationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListFederationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFederationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFederationsResponseBody) SetFederations(v []*ListFederationsResponseBodyFederations) *ListFederationsResponseBody {
	s.Federations = v
	return s
}

func (s *ListFederationsResponseBody) SetMaxResults(v int32) *ListFederationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListFederationsResponseBody) SetNextToken(v string) *ListFederationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFederationsResponseBody) SetRequestId(v string) *ListFederationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFederationsResponseBody) Validate() error {
	if s.Federations != nil {
		for _, item := range s.Federations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFederationsResponseBodyFederations struct {
	FederationId  *string `json:"FederationId,omitempty" xml:"FederationId,omitempty"`
	FileSystemIds *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
}

func (s ListFederationsResponseBodyFederations) String() string {
	return dara.Prettify(s)
}

func (s ListFederationsResponseBodyFederations) GoString() string {
	return s.String()
}

func (s *ListFederationsResponseBodyFederations) GetFederationId() *string {
	return s.FederationId
}

func (s *ListFederationsResponseBodyFederations) GetFileSystemIds() *string {
	return s.FileSystemIds
}

func (s *ListFederationsResponseBodyFederations) SetFederationId(v string) *ListFederationsResponseBodyFederations {
	s.FederationId = &v
	return s
}

func (s *ListFederationsResponseBodyFederations) SetFileSystemIds(v string) *ListFederationsResponseBodyFederations {
	s.FileSystemIds = &v
	return s
}

func (s *ListFederationsResponseBodyFederations) Validate() error {
	return dara.Validate(s)
}
