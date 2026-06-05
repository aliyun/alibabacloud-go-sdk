// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListAppPluginsRequest
	GetBizId() *string
	SetMaxResults(v int32) *ListAppPluginsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppPluginsRequest
	GetNextToken() *string
	SetPhase(v string) *ListAppPluginsRequest
	GetPhase() *string
	SetPlatform(v string) *ListAppPluginsRequest
	GetPlatform() *string
}

type ListAppPluginsRequest struct {
	// example:
	//
	// WD20250820143531000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// http_whitelist
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// linux/amd64
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s ListAppPluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppPluginsRequest) GoString() string {
	return s.String()
}

func (s *ListAppPluginsRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListAppPluginsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppPluginsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppPluginsRequest) GetPhase() *string {
	return s.Phase
}

func (s *ListAppPluginsRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListAppPluginsRequest) SetBizId(v string) *ListAppPluginsRequest {
	s.BizId = &v
	return s
}

func (s *ListAppPluginsRequest) SetMaxResults(v int32) *ListAppPluginsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppPluginsRequest) SetNextToken(v string) *ListAppPluginsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppPluginsRequest) SetPhase(v string) *ListAppPluginsRequest {
	s.Phase = &v
	return s
}

func (s *ListAppPluginsRequest) SetPlatform(v string) *ListAppPluginsRequest {
	s.Platform = &v
	return s
}

func (s *ListAppPluginsRequest) Validate() error {
	return dara.Validate(s)
}
