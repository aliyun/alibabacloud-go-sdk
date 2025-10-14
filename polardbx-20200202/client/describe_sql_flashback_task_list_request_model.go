// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlFlashbackTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolardbxInstanceId(v string) *DescribeSqlFlashbackTaskListRequest
	GetPolardbxInstanceId() *string
	SetRegionId(v string) *DescribeSqlFlashbackTaskListRequest
	GetRegionId() *string
}

type DescribeSqlFlashbackTaskListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-**********
	PolardbxInstanceId *string `json:"PolardbxInstanceId,omitempty" xml:"PolardbxInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSqlFlashbackTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlFlashbackTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlFlashbackTaskListRequest) GetPolardbxInstanceId() *string {
	return s.PolardbxInstanceId
}

func (s *DescribeSqlFlashbackTaskListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSqlFlashbackTaskListRequest) SetPolardbxInstanceId(v string) *DescribeSqlFlashbackTaskListRequest {
	s.PolardbxInstanceId = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListRequest) SetRegionId(v string) *DescribeSqlFlashbackTaskListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSqlFlashbackTaskListRequest) Validate() error {
	return dara.Validate(s)
}
