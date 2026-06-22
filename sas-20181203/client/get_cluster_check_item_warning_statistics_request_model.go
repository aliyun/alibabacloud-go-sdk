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
	// The ID of the container cluster that you want to query.
	//
	// > Call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/182997.html) operation to obtain this parameter.
	//
	// example:
	//
	// cd49575861a3044d49c954e4b3911****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The UUID of the server.
	//
	// > Call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// ae1527a9-2308-46ab-b10a-48ae7ff7****
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
