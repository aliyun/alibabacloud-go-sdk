// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiDestinationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiDestinationNamePrefix(v string) *ListApiDestinationsRequest
	GetApiDestinationNamePrefix() *string
	SetConnectionName(v string) *ListApiDestinationsRequest
	GetConnectionName() *string
	SetMaxResults(v int64) *ListApiDestinationsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListApiDestinationsRequest
	GetNextToken() *string
}

type ListApiDestinationsRequest struct {
	// The prefix of the API destination name.
	//
	// example:
	//
	// api-demo
	ApiDestinationNamePrefix *string `json:"ApiDestinationNamePrefix,omitempty" xml:"ApiDestinationNamePrefix,omitempty"`
	// The connection name.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The maximum number of entries to be returned in a call. You can use this parameter and NextToken to implement paging.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// If you set Limit and excess return values exist, this parameter is returned.
	//
	// 	- Default value: 0.
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListApiDestinationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiDestinationsRequest) GoString() string {
	return s.String()
}

func (s *ListApiDestinationsRequest) GetApiDestinationNamePrefix() *string {
	return s.ApiDestinationNamePrefix
}

func (s *ListApiDestinationsRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *ListApiDestinationsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListApiDestinationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiDestinationsRequest) SetApiDestinationNamePrefix(v string) *ListApiDestinationsRequest {
	s.ApiDestinationNamePrefix = &v
	return s
}

func (s *ListApiDestinationsRequest) SetConnectionName(v string) *ListApiDestinationsRequest {
	s.ConnectionName = &v
	return s
}

func (s *ListApiDestinationsRequest) SetMaxResults(v int64) *ListApiDestinationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApiDestinationsRequest) SetNextToken(v string) *ListApiDestinationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListApiDestinationsRequest) Validate() error {
	return dara.Validate(s)
}
