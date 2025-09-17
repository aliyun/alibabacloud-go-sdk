// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeName(v string) *ListDataSourcesRequest
	GetAttributeName() *string
	SetFilter(v map[string]interface{}) *ListDataSourcesRequest
	GetFilter() map[string]interface{}
}

type ListDataSourcesRequest struct {
	// The name of the property. RegionId is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// RegionId
	AttributeName *string `json:"attributeName,omitempty" xml:"attributeName,omitempty"`
	// The filter conditions. JSON format:{"key1":"value1"}.
	Filter map[string]interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
}

func (s ListDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesRequest) GetAttributeName() *string {
	return s.AttributeName
}

func (s *ListDataSourcesRequest) GetFilter() map[string]interface{} {
	return s.Filter
}

func (s *ListDataSourcesRequest) SetAttributeName(v string) *ListDataSourcesRequest {
	s.AttributeName = &v
	return s
}

func (s *ListDataSourcesRequest) SetFilter(v map[string]interface{}) *ListDataSourcesRequest {
	s.Filter = v
	return s
}

func (s *ListDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
