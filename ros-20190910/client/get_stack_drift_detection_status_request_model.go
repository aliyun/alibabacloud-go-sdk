// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackDriftDetectionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriftDetectionId(v string) *GetStackDriftDetectionStatusRequest
	GetDriftDetectionId() *string
	SetRegionId(v string) *GetStackDriftDetectionStatusRequest
	GetRegionId() *string
}

type GetStackDriftDetectionStatusRequest struct {
	// The ID of the drift detection operation.
	//
	// You can call the [ListStackResourceDrifts](https://help.aliyun.com/document_detail/155098.html) operation to obtain the ID of the drift detection operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// a7044f0d-6f2e-4128-a307-4524ef88****
	DriftDetectionId *string `json:"DriftDetectionId,omitempty" xml:"DriftDetectionId,omitempty"`
	// The region ID of the stack to be detected for drift.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetStackDriftDetectionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackDriftDetectionStatusRequest) GoString() string {
	return s.String()
}

func (s *GetStackDriftDetectionStatusRequest) GetDriftDetectionId() *string {
	return s.DriftDetectionId
}

func (s *GetStackDriftDetectionStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStackDriftDetectionStatusRequest) SetDriftDetectionId(v string) *GetStackDriftDetectionStatusRequest {
	s.DriftDetectionId = &v
	return s
}

func (s *GetStackDriftDetectionStatusRequest) SetRegionId(v string) *GetStackDriftDetectionStatusRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackDriftDetectionStatusRequest) Validate() error {
	return dara.Validate(s)
}
