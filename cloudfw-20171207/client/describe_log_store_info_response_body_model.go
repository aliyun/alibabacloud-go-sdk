// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogStoreName(v string) *DescribeLogStoreInfoResponseBody
	GetLogStoreName() *string
	SetProjectName(v string) *DescribeLogStoreInfoResponseBody
	GetProjectName() *string
	SetQuota(v int64) *DescribeLogStoreInfoResponseBody
	GetQuota() *int64
	SetRegionId(v string) *DescribeLogStoreInfoResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeLogStoreInfoResponseBody
	GetRequestId() *string
	SetTtl(v int32) *DescribeLogStoreInfoResponseBody
	GetTtl() *int32
	SetUsed(v int64) *DescribeLogStoreInfoResponseBody
	GetUsed() *int64
}

type DescribeLogStoreInfoResponseBody struct {
	// The name of the SLS LogStore in the log service.
	//
	// example:
	//
	// xxx-logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The Project name of the log service.
	//
	// example:
	//
	// project-xxx-cn-hangzhou
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Available log storage capacity. Unit: Byte.
	//
	// example:
	//
	// 50000000
	Quota *int64 `json:"Quota,omitempty" xml:"Quota,omitempty"`
	// The region ID for log delivery.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of this request.
	//
	// example:
	//
	// C6C3B72B********E95FB0A161
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Log storage duration. Unit: days.
	//
	// example:
	//
	// 20
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// Used storage capacity. Unit: Byte.
	//
	// > The statistics of the log service have a delay of approximately two hours.
	//
	// example:
	//
	// 0
	Used *int64 `json:"Used,omitempty" xml:"Used,omitempty"`
}

func (s DescribeLogStoreInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreInfoResponseBody) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeLogStoreInfoResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeLogStoreInfoResponseBody) GetQuota() *int64 {
	return s.Quota
}

func (s *DescribeLogStoreInfoResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLogStoreInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogStoreInfoResponseBody) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeLogStoreInfoResponseBody) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeLogStoreInfoResponseBody) SetLogStoreName(v string) *DescribeLogStoreInfoResponseBody {
	s.LogStoreName = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetProjectName(v string) *DescribeLogStoreInfoResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetQuota(v int64) *DescribeLogStoreInfoResponseBody {
	s.Quota = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetRegionId(v string) *DescribeLogStoreInfoResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetRequestId(v string) *DescribeLogStoreInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetTtl(v int32) *DescribeLogStoreInfoResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) SetUsed(v int64) *DescribeLogStoreInfoResponseBody {
	s.Used = &v
	return s
}

func (s *DescribeLogStoreInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
