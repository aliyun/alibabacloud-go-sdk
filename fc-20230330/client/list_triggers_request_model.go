// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTriggersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListTriggersRequest
	GetLimit() *int32
	SetNextToken(v string) *ListTriggersRequest
	GetNextToken() *string
	SetPrefix(v string) *ListTriggersRequest
	GetPrefix() *string
}

type ListTriggersRequest struct {
	// The number of triggers returned.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// MTIzNCNhYmM=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The trigger name prefix.
	//
	// example:
	//
	// my-trigger
	Prefix *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListTriggersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTriggersRequest) GoString() string {
	return s.String()
}

func (s *ListTriggersRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListTriggersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTriggersRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *ListTriggersRequest) SetLimit(v int32) *ListTriggersRequest {
	s.Limit = &v
	return s
}

func (s *ListTriggersRequest) SetNextToken(v string) *ListTriggersRequest {
	s.NextToken = &v
	return s
}

func (s *ListTriggersRequest) SetPrefix(v string) *ListTriggersRequest {
	s.Prefix = &v
	return s
}

func (s *ListTriggersRequest) Validate() error {
	return dara.Validate(s)
}
