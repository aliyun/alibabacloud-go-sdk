// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterConfiguration(v string) *CreateFilterRequest
	GetFilterConfiguration() *string
	SetFilterName(v string) *CreateFilterRequest
	GetFilterName() *string
}

type CreateFilterRequest struct {
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

func (s CreateFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFilterRequest) GoString() string {
	return s.String()
}

func (s *CreateFilterRequest) GetFilterConfiguration() *string {
	return s.FilterConfiguration
}

func (s *CreateFilterRequest) GetFilterName() *string {
	return s.FilterName
}

func (s *CreateFilterRequest) SetFilterConfiguration(v string) *CreateFilterRequest {
	s.FilterConfiguration = &v
	return s
}

func (s *CreateFilterRequest) SetFilterName(v string) *CreateFilterRequest {
	s.FilterName = &v
	return s
}

func (s *CreateFilterRequest) Validate() error {
	return dara.Validate(s)
}
