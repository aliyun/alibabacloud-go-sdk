// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListForwardStrategiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *ListForwardStrategiesRequest
	GetCurrentPage() *string
	SetDestinationIds(v []*string) *ListForwardStrategiesRequest
	GetDestinationIds() []*string
	SetDestinationType(v string) *ListForwardStrategiesRequest
	GetDestinationType() *string
	SetForwardIds(v []*string) *ListForwardStrategiesRequest
	GetForwardIds() []*string
	SetName(v string) *ListForwardStrategiesRequest
	GetName() *string
	SetPageSize(v string) *ListForwardStrategiesRequest
	GetPageSize() *string
}

type ListForwardStrategiesRequest struct {
	// example:
	//
	// 1
	CurrentPage    *string   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DestinationIds []*string `json:"DestinationIds,omitempty" xml:"DestinationIds,omitempty" type:"Repeated"`
	// example:
	//
	// Connector
	DestinationType *string   `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	ForwardIds      []*string `json:"ForwardIds,omitempty" xml:"ForwardIds,omitempty" type:"Repeated"`
	// example:
	//
	// acs_rand_str_acs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListForwardStrategiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListForwardStrategiesRequest) GoString() string {
	return s.String()
}

func (s *ListForwardStrategiesRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *ListForwardStrategiesRequest) GetDestinationIds() []*string {
	return s.DestinationIds
}

func (s *ListForwardStrategiesRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *ListForwardStrategiesRequest) GetForwardIds() []*string {
	return s.ForwardIds
}

func (s *ListForwardStrategiesRequest) GetName() *string {
	return s.Name
}

func (s *ListForwardStrategiesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListForwardStrategiesRequest) SetCurrentPage(v string) *ListForwardStrategiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListForwardStrategiesRequest) SetDestinationIds(v []*string) *ListForwardStrategiesRequest {
	s.DestinationIds = v
	return s
}

func (s *ListForwardStrategiesRequest) SetDestinationType(v string) *ListForwardStrategiesRequest {
	s.DestinationType = &v
	return s
}

func (s *ListForwardStrategiesRequest) SetForwardIds(v []*string) *ListForwardStrategiesRequest {
	s.ForwardIds = v
	return s
}

func (s *ListForwardStrategiesRequest) SetName(v string) *ListForwardStrategiesRequest {
	s.Name = &v
	return s
}

func (s *ListForwardStrategiesRequest) SetPageSize(v string) *ListForwardStrategiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListForwardStrategiesRequest) Validate() error {
	return dara.Validate(s)
}
