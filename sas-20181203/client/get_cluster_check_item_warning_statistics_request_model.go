// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterCheckItemWarningStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetClusterCheckItemWarningStatisticsRequest
	GetClusterId() *string
	SetUuid(v string) *GetClusterCheckItemWarningStatisticsRequest
	GetUuid() *string
}

type GetClusterCheckItemWarningStatisticsRequest struct {
	// The ID of the container cluster.
	//
	// >  You can call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/182997.html) operation to query the IDs of container clusters.
	//
	// example:
	//
	// cd49575861a3044d49c954e4b3911****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The server UUID.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the server UUID.
	//
	// example:
	//
	// e4af3620-6895-4e2f-a641-a9d8fb53****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetClusterCheckItemWarningStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterCheckItemWarningStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetClusterCheckItemWarningStatisticsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterCheckItemWarningStatisticsRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetClusterCheckItemWarningStatisticsRequest) SetClusterId(v string) *GetClusterCheckItemWarningStatisticsRequest {
	s.ClusterId = &v
	return s
}

func (s *GetClusterCheckItemWarningStatisticsRequest) SetUuid(v string) *GetClusterCheckItemWarningStatisticsRequest {
	s.Uuid = &v
	return s
}

func (s *GetClusterCheckItemWarningStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
