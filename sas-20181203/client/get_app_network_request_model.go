// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetAppNetworkRequest
	GetClusterId() *string
	SetEndTime(v int64) *GetAppNetworkRequest
	GetEndTime() *int64
	SetStartTime(v int64) *GetAppNetworkRequest
	GetStartTime() *int64
}

type GetAppNetworkRequest struct {
	// The ID of the cluster to which the container belongs.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of container clusters.
	//
	// example:
	//
	// cf77xxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end timestamp of the query. Unit: milliseconds.
	//
	// > The days between the start timestamp and the end timestamp cannot exceed **seven*	- days.
	//
	// example:
	//
	// 1650470399999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start timestamp of the query. Unit: milliseconds.
	//
	// > The days between the start timestamp and the end timestamp cannot exceed **seven*	- days.
	//
	// example:
	//
	// 1649260800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAppNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppNetworkRequest) GoString() string {
	return s.String()
}

func (s *GetAppNetworkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetAppNetworkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetAppNetworkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetAppNetworkRequest) SetClusterId(v string) *GetAppNetworkRequest {
	s.ClusterId = &v
	return s
}

func (s *GetAppNetworkRequest) SetEndTime(v int64) *GetAppNetworkRequest {
	s.EndTime = &v
	return s
}

func (s *GetAppNetworkRequest) SetStartTime(v int64) *GetAppNetworkRequest {
	s.StartTime = &v
	return s
}

func (s *GetAppNetworkRequest) Validate() error {
	return dara.Validate(s)
}
