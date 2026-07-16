// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAWSRegionInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListAWSRegionInfosRequest
	GetName() *string
}

type ListAWSRegionInfosRequest struct {
	// The name of the metadata to query.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAWSRegionInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAWSRegionInfosRequest) GoString() string {
	return s.String()
}

func (s *ListAWSRegionInfosRequest) GetName() *string {
	return s.Name
}

func (s *ListAWSRegionInfosRequest) SetName(v string) *ListAWSRegionInfosRequest {
	s.Name = &v
	return s
}

func (s *ListAWSRegionInfosRequest) Validate() error {
	return dara.Validate(s)
}
