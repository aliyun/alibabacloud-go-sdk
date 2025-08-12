// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteCustomMetricRequest
	GetGroupId() *string
	SetMd5(v string) *DeleteCustomMetricRequest
	GetMd5() *string
	SetMetricName(v string) *DeleteCustomMetricRequest
	GetMetricName() *string
	SetRegionId(v string) *DeleteCustomMetricRequest
	GetRegionId() *string
	SetUUID(v string) *DeleteCustomMetricRequest
	GetUUID() *string
}

type DeleteCustomMetricRequest struct {
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3607****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The MD5 value of the HTTP request body. The MD5 value is a 128-bit hash value used to verify the uniqueness of the reported monitoring data.
	//
	// >  `Md5` is returned when you query the reported monitoring data of a metric.
	//
	// example:
	//
	// 38796C8CFFEB8F89BB2A626C7BD7****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the metric.
	//
	// This parameter is required.
	//
	// example:
	//
	// AdvanceCredit
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request for reporting monitoring data.
	//
	// >  `UUID` is returned when you query the reported monitoring data of a metric. We recommend that you specify the `Md5` parameter.
	//
	// example:
	//
	// 5497633c-66c5-4eae-abaa-89db5adb****
	UUID *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s DeleteCustomMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomMetricRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomMetricRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteCustomMetricRequest) GetMd5() *string {
	return s.Md5
}

func (s *DeleteCustomMetricRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DeleteCustomMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomMetricRequest) GetUUID() *string {
	return s.UUID
}

func (s *DeleteCustomMetricRequest) SetGroupId(v string) *DeleteCustomMetricRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteCustomMetricRequest) SetMd5(v string) *DeleteCustomMetricRequest {
	s.Md5 = &v
	return s
}

func (s *DeleteCustomMetricRequest) SetMetricName(v string) *DeleteCustomMetricRequest {
	s.MetricName = &v
	return s
}

func (s *DeleteCustomMetricRequest) SetRegionId(v string) *DeleteCustomMetricRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomMetricRequest) SetUUID(v string) *DeleteCustomMetricRequest {
	s.UUID = &v
	return s
}

func (s *DeleteCustomMetricRequest) Validate() error {
	return dara.Validate(s)
}
