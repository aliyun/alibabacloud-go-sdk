// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAuthConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthType(v string) *ListUserAuthConfigsRequest
	GetAuthType() *string
	SetConnectorId(v string) *ListUserAuthConfigsRequest
	GetConnectorId() *string
	SetConnectorVersion(v string) *ListUserAuthConfigsRequest
	GetConnectorVersion() *string
	SetFilter(v []*ListUserAuthConfigsRequestFilter) *ListUserAuthConfigsRequest
	GetFilter() []*ListUserAuthConfigsRequestFilter
	SetMaxResults(v string) *ListUserAuthConfigsRequest
	GetMaxResults() *string
	SetNextToken(v string) *ListUserAuthConfigsRequest
	GetNextToken() *string
}

type ListUserAuthConfigsRequest struct {
	// example:
	//
	// QQBotAccessToken
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// connector-88d2c03da8c9410e8a91
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// 6
	ConnectorVersion *string                             `json:"ConnectorVersion,omitempty" xml:"ConnectorVersion,omitempty"`
	Filter           []*ListUserAuthConfigsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAVY3rYiv9VoUJQSiCitgjgQu5rMgGgvUoNWg8LykhA85j8bgHiGAwZWnCMJPepC+WWc0DK5hx1qIycMHVWP2AjQ=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListUserAuthConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListUserAuthConfigsRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ListUserAuthConfigsRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *ListUserAuthConfigsRequest) GetConnectorVersion() *string {
	return s.ConnectorVersion
}

func (s *ListUserAuthConfigsRequest) GetFilter() []*ListUserAuthConfigsRequestFilter {
	return s.Filter
}

func (s *ListUserAuthConfigsRequest) GetMaxResults() *string {
	return s.MaxResults
}

func (s *ListUserAuthConfigsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserAuthConfigsRequest) SetAuthType(v string) *ListUserAuthConfigsRequest {
	s.AuthType = &v
	return s
}

func (s *ListUserAuthConfigsRequest) SetConnectorId(v string) *ListUserAuthConfigsRequest {
	s.ConnectorId = &v
	return s
}

func (s *ListUserAuthConfigsRequest) SetConnectorVersion(v string) *ListUserAuthConfigsRequest {
	s.ConnectorVersion = &v
	return s
}

func (s *ListUserAuthConfigsRequest) SetFilter(v []*ListUserAuthConfigsRequestFilter) *ListUserAuthConfigsRequest {
	s.Filter = v
	return s
}

func (s *ListUserAuthConfigsRequest) SetMaxResults(v string) *ListUserAuthConfigsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserAuthConfigsRequest) SetNextToken(v string) *ListUserAuthConfigsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserAuthConfigsRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserAuthConfigsRequestFilter struct {
	// example:
	//
	// AuthConfigName
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListUserAuthConfigsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListUserAuthConfigsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListUserAuthConfigsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListUserAuthConfigsRequestFilter) GetValues() []*string {
	return s.Values
}

func (s *ListUserAuthConfigsRequestFilter) SetName(v string) *ListUserAuthConfigsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListUserAuthConfigsRequestFilter) SetValues(v []*string) *ListUserAuthConfigsRequestFilter {
	s.Values = v
	return s
}

func (s *ListUserAuthConfigsRequestFilter) Validate() error {
	return dara.Validate(s)
}
