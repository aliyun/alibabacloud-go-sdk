// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeContainerStatisticsRequest
	GetClusterId() *string
}

type DescribeContainerStatisticsRequest struct {
	// The ID of the specified container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to obtain the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Cccfd68c474454665ace07efce924****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeContainerStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerStatisticsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerStatisticsRequest) SetClusterId(v string) *DescribeContainerStatisticsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
