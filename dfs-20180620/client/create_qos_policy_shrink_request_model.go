// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateQosPolicyShrinkRequest
	GetDescription() *string
	SetFederationId(v string) *CreateQosPolicyShrinkRequest
	GetFederationId() *string
	SetFileSystemId(v string) *CreateQosPolicyShrinkRequest
	GetFileSystemId() *string
	SetFlowIdsShrink(v string) *CreateQosPolicyShrinkRequest
	GetFlowIdsShrink() *string
	SetInputRegionId(v string) *CreateQosPolicyShrinkRequest
	GetInputRegionId() *string
	SetMaxIOBandWidth(v int64) *CreateQosPolicyShrinkRequest
	GetMaxIOBandWidth() *int64
	SetMaxIOps(v int64) *CreateQosPolicyShrinkRequest
	GetMaxIOps() *int64
	SetMaxMetaQps(v int64) *CreateQosPolicyShrinkRequest
	GetMaxMetaQps() *int64
	SetReqTagsShrink(v string) *CreateQosPolicyShrinkRequest
	GetReqTagsShrink() *string
	SetZoneIdsShrink(v string) *CreateQosPolicyShrinkRequest
	GetZoneIdsShrink() *string
}

type CreateQosPolicyShrinkRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FederationId  *string `json:"FederationId,omitempty" xml:"FederationId,omitempty"`
	FileSystemId  *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FlowIdsShrink *string `json:"FlowIds,omitempty" xml:"FlowIds,omitempty"`
	// This parameter is required.
	InputRegionId  *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	MaxIOBandWidth *int64  `json:"MaxIOBandWidth,omitempty" xml:"MaxIOBandWidth,omitempty"`
	MaxIOps        *int64  `json:"MaxIOps,omitempty" xml:"MaxIOps,omitempty"`
	MaxMetaQps     *int64  `json:"MaxMetaQps,omitempty" xml:"MaxMetaQps,omitempty"`
	ReqTagsShrink  *string `json:"ReqTags,omitempty" xml:"ReqTags,omitempty"`
	ZoneIdsShrink  *string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
}

func (s CreateQosPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQosPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateQosPolicyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateQosPolicyShrinkRequest) GetFederationId() *string {
	return s.FederationId
}

func (s *CreateQosPolicyShrinkRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateQosPolicyShrinkRequest) GetFlowIdsShrink() *string {
	return s.FlowIdsShrink
}

func (s *CreateQosPolicyShrinkRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *CreateQosPolicyShrinkRequest) GetMaxIOBandWidth() *int64 {
	return s.MaxIOBandWidth
}

func (s *CreateQosPolicyShrinkRequest) GetMaxIOps() *int64 {
	return s.MaxIOps
}

func (s *CreateQosPolicyShrinkRequest) GetMaxMetaQps() *int64 {
	return s.MaxMetaQps
}

func (s *CreateQosPolicyShrinkRequest) GetReqTagsShrink() *string {
	return s.ReqTagsShrink
}

func (s *CreateQosPolicyShrinkRequest) GetZoneIdsShrink() *string {
	return s.ZoneIdsShrink
}

func (s *CreateQosPolicyShrinkRequest) SetDescription(v string) *CreateQosPolicyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateQosPolicyShrinkRequest) SetFederationId(v string) *CreateQosPolicyShrinkRequest {
	s.FederationId = &v
	return s
}

func (s *CreateQosPolicyShrinkRequest) SetFileSystemId(v string) *CreateQosPolicyShrinkRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateQosPolicyShrinkRequest) SetFlowIdsShrink(v string) *CreateQosPolicyShrinkRequest {
	s.FlowIdsShrink = &v
	return s
}

func (s *CreateQosPolicyShrinkRequest) SetInputRegionId(v string) *CreateQosPolicyShrinkRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateQosPolicyShrinkRequest) SetMaxIOBandWidth(v int64) *CreateQosPolicyShrinkRequest {
	s.MaxIOBandWidth = &v
	return s
}

func (s *CreateQosPolicyShrinkRequest) SetMaxIOps(v int64) *CreateQosPolicyShrinkRequest {
	s.MaxIOps = &v
	return s
}

func (s *CreateQosPolicyShrinkRequest) SetMaxMetaQps(v int64) *CreateQosPolicyShrinkRequest {
	s.MaxMetaQps = &v
	return s
}

func (s *CreateQosPolicyShrinkRequest) SetReqTagsShrink(v string) *CreateQosPolicyShrinkRequest {
	s.ReqTagsShrink = &v
	return s
}

func (s *CreateQosPolicyShrinkRequest) SetZoneIdsShrink(v string) *CreateQosPolicyShrinkRequest {
	s.ZoneIdsShrink = &v
	return s
}

func (s *CreateQosPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
