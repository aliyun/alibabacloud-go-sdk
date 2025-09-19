// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterSuspEventStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetClusterSuspEventStatisticsRequest
	GetClusterId() *string
	SetFrom(v string) *GetClusterSuspEventStatisticsRequest
	GetFrom() *string
}

type GetClusterSuspEventStatisticsRequest struct {
	// The ID of the container cluster.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of container clusters.
	//
	// example:
	//
	// c6094b964bfc145fe9e418c869e7e****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the request source. Set the value to sas.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
}

func (s GetClusterSuspEventStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterSuspEventStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetClusterSuspEventStatisticsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterSuspEventStatisticsRequest) GetFrom() *string {
	return s.From
}

func (s *GetClusterSuspEventStatisticsRequest) SetClusterId(v string) *GetClusterSuspEventStatisticsRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterSuspEventStatisticsRequest) SetFrom(v string) *GetClusterSuspEventStatisticsRequest {
	s.From = &v
	return s
}

func (s *GetClusterSuspEventStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
