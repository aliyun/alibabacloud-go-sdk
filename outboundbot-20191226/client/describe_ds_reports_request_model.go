// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDsReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDsReportsRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *DescribeDsReportsRequest
	GetJobGroupId() *string
}

type DescribeDsReportsRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// a4274627-265f-4e14-b2d6-4ee7d4f8593e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Task group ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 9eb8fe4f-b286-4834-9688-2c9b171e223e
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
}

func (s DescribeDsReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDsReportsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDsReportsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDsReportsRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *DescribeDsReportsRequest) SetInstanceId(v string) *DescribeDsReportsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDsReportsRequest) SetJobGroupId(v string) *DescribeDsReportsRequest {
	s.JobGroupId = &v
	return s
}

func (s *DescribeDsReportsRequest) Validate() error {
	return dara.Validate(s)
}
