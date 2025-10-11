// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterConfiguration(v string) *UpdateFilterRequest
	GetFilterConfiguration() *string
	SetFilterName(v string) *UpdateFilterRequest
	GetFilterName() *string
}

type UpdateFilterRequest struct {
	// The configurations of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "regions": [],
	//
	//   "tagFilters": [
	//
	//     [{ "type": "notContainTagKey", "tagKey": "xxx", "tagValue": "" }],
	//
	//     [{ "tagKey": "xxx", "tagValue": "xxx" }]
	//
	//   ],
	//
	//   "resourceTypes": [
	//
	//     "ACS::ECS::AutoSnapshotPolicy"
	//
	//   ]
	//
	// }
	FilterConfiguration *string `json:"FilterConfiguration,omitempty" xml:"FilterConfiguration,omitempty"`
	// The name of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s UpdateFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFilterRequest) GoString() string {
	return s.String()
}

func (s *UpdateFilterRequest) GetFilterConfiguration() *string {
	return s.FilterConfiguration
}

func (s *UpdateFilterRequest) GetFilterName() *string {
	return s.FilterName
}

func (s *UpdateFilterRequest) SetFilterConfiguration(v string) *UpdateFilterRequest {
	s.FilterConfiguration = &v
	return s
}

func (s *UpdateFilterRequest) SetFilterName(v string) *UpdateFilterRequest {
	s.FilterName = &v
	return s
}

func (s *UpdateFilterRequest) Validate() error {
	return dara.Validate(s)
}
