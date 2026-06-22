// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetImageBuildRiskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUuids(v string) *SetImageBuildRiskStatusRequest
	GetImageUuids() *string
	SetRiskKey(v string) *SetImageBuildRiskStatusRequest
	GetRiskKey() *string
	SetStatus(v int32) *SetImageBuildRiskStatusRequest
	GetStatus() *int32
}

type SetImageBuildRiskStatusRequest struct {
	// The image UUIDs. Separate multiple UUIDs with commas (,).
	//
	// >Call the [DescribeImageInstances](~~DescribeImageInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// f382fccd88b94c5c8c864def681*****,ac32fccd88b94c5c8c864def681*****
	ImageUuids *string `json:"ImageUuids,omitempty" xml:"ImageUuids,omitempty"`
	// The risk keyword.
	//
	// example:
	//
	// risk.type
	RiskKey *string `json:"RiskKey,omitempty" xml:"RiskKey,omitempty"`
	// The status. Valid values:
	//
	// - **0**: Unhandled.
	//
	// - **1**: Ignored.
	//
	// - **2**: False positive.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetImageBuildRiskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetImageBuildRiskStatusRequest) GoString() string {
	return s.String()
}

func (s *SetImageBuildRiskStatusRequest) GetImageUuids() *string {
	return s.ImageUuids
}

func (s *SetImageBuildRiskStatusRequest) GetRiskKey() *string {
	return s.RiskKey
}

func (s *SetImageBuildRiskStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *SetImageBuildRiskStatusRequest) SetImageUuids(v string) *SetImageBuildRiskStatusRequest {
	s.ImageUuids = &v
	return s
}

func (s *SetImageBuildRiskStatusRequest) SetRiskKey(v string) *SetImageBuildRiskStatusRequest {
	s.RiskKey = &v
	return s
}

func (s *SetImageBuildRiskStatusRequest) SetStatus(v int32) *SetImageBuildRiskStatusRequest {
	s.Status = &v
	return s
}

func (s *SetImageBuildRiskStatusRequest) Validate() error {
	return dara.Validate(s)
}
