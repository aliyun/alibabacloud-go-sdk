// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertStrategiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int32) *ListAlertStrategiesRequest
	GetCurrent() *int32
	SetMaxResults(v int32) *ListAlertStrategiesRequest
	GetMaxResults() *int32
	SetName(v string) *ListAlertStrategiesRequest
	GetName() *string
	SetNextToken(v string) *ListAlertStrategiesRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListAlertStrategiesRequest
	GetPageSize() *int32
}

type ListAlertStrategiesRequest struct {
	// example:
	//
	// 1
	Current *int32 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// strategy1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// c2f78a783f49457caba6bace6f6f79e4
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListAlertStrategiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertStrategiesRequest) GoString() string {
	return s.String()
}

func (s *ListAlertStrategiesRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAlertStrategiesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertStrategiesRequest) GetName() *string {
	return s.Name
}

func (s *ListAlertStrategiesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertStrategiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertStrategiesRequest) SetCurrent(v int32) *ListAlertStrategiesRequest {
	s.Current = &v
	return s
}

func (s *ListAlertStrategiesRequest) SetMaxResults(v int32) *ListAlertStrategiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAlertStrategiesRequest) SetName(v string) *ListAlertStrategiesRequest {
	s.Name = &v
	return s
}

func (s *ListAlertStrategiesRequest) SetNextToken(v string) *ListAlertStrategiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAlertStrategiesRequest) SetPageSize(v int32) *ListAlertStrategiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertStrategiesRequest) Validate() error {
	return dara.Validate(s)
}
