// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *GetProductListRequest
	GetApiKey() *string
	SetConfigVersion(v string) *GetProductListRequest
	GetConfigVersion() *string
	SetMaxResults(v int32) *GetProductListRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetProductListRequest
	GetNextToken() *string
	SetSessionId(v string) *GetProductListRequest
	GetSessionId() *string
	SetType(v string) *GetProductListRequest
	GetType() *string
}

type GetProductListRequest struct {
	// This parameter is required.
	ApiKey        *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	ConfigVersion *string `json:"ConfigVersion,omitempty" xml:"ConfigVersion,omitempty"`
	MaxResults    *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	SessionId     *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetProductListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductListRequest) GoString() string {
	return s.String()
}

func (s *GetProductListRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *GetProductListRequest) GetConfigVersion() *string {
	return s.ConfigVersion
}

func (s *GetProductListRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetProductListRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetProductListRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetProductListRequest) GetType() *string {
	return s.Type
}

func (s *GetProductListRequest) SetApiKey(v string) *GetProductListRequest {
	s.ApiKey = &v
	return s
}

func (s *GetProductListRequest) SetConfigVersion(v string) *GetProductListRequest {
	s.ConfigVersion = &v
	return s
}

func (s *GetProductListRequest) SetMaxResults(v int32) *GetProductListRequest {
	s.MaxResults = &v
	return s
}

func (s *GetProductListRequest) SetNextToken(v string) *GetProductListRequest {
	s.NextToken = &v
	return s
}

func (s *GetProductListRequest) SetSessionId(v string) *GetProductListRequest {
	s.SessionId = &v
	return s
}

func (s *GetProductListRequest) SetType(v string) *GetProductListRequest {
	s.Type = &v
	return s
}

func (s *GetProductListRequest) Validate() error {
	return dara.Validate(s)
}
