// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateQosPolicyRequest
	GetDescription() *string
	SetFederationId(v string) *CreateQosPolicyRequest
	GetFederationId() *string
	SetFileSystemId(v string) *CreateQosPolicyRequest
	GetFileSystemId() *string
	SetFlowIds(v []*int32) *CreateQosPolicyRequest
	GetFlowIds() []*int32
	SetInputRegionId(v string) *CreateQosPolicyRequest
	GetInputRegionId() *string
	SetMaxIOBandWidth(v int64) *CreateQosPolicyRequest
	GetMaxIOBandWidth() *int64
	SetMaxIOps(v int64) *CreateQosPolicyRequest
	GetMaxIOps() *int64
	SetMaxMetaQps(v int64) *CreateQosPolicyRequest
	GetMaxMetaQps() *int64
	SetReqTags(v []*string) *CreateQosPolicyRequest
	GetReqTags() []*string
	SetZoneIds(v []*string) *CreateQosPolicyRequest
	GetZoneIds() []*string
}

type CreateQosPolicyRequest struct {
	Description  *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	FederationId *string  `json:"FederationId,omitempty" xml:"FederationId,omitempty"`
	FileSystemId *string  `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	FlowIds      []*int32 `json:"FlowIds,omitempty" xml:"FlowIds,omitempty" type:"Repeated"`
	// This parameter is required.
	InputRegionId  *string   `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	MaxIOBandWidth *int64    `json:"MaxIOBandWidth,omitempty" xml:"MaxIOBandWidth,omitempty"`
	MaxIOps        *int64    `json:"MaxIOps,omitempty" xml:"MaxIOps,omitempty"`
	MaxMetaQps     *int64    `json:"MaxMetaQps,omitempty" xml:"MaxMetaQps,omitempty"`
	ReqTags        []*string `json:"ReqTags,omitempty" xml:"ReqTags,omitempty" type:"Repeated"`
	ZoneIds        []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
}

func (s CreateQosPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQosPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateQosPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateQosPolicyRequest) GetFederationId() *string {
	return s.FederationId
}

func (s *CreateQosPolicyRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateQosPolicyRequest) GetFlowIds() []*int32 {
	return s.FlowIds
}

func (s *CreateQosPolicyRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *CreateQosPolicyRequest) GetMaxIOBandWidth() *int64 {
	return s.MaxIOBandWidth
}

func (s *CreateQosPolicyRequest) GetMaxIOps() *int64 {
	return s.MaxIOps
}

func (s *CreateQosPolicyRequest) GetMaxMetaQps() *int64 {
	return s.MaxMetaQps
}

func (s *CreateQosPolicyRequest) GetReqTags() []*string {
	return s.ReqTags
}

func (s *CreateQosPolicyRequest) GetZoneIds() []*string {
	return s.ZoneIds
}

func (s *CreateQosPolicyRequest) SetDescription(v string) *CreateQosPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateQosPolicyRequest) SetFederationId(v string) *CreateQosPolicyRequest {
	s.FederationId = &v
	return s
}

func (s *CreateQosPolicyRequest) SetFileSystemId(v string) *CreateQosPolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateQosPolicyRequest) SetFlowIds(v []*int32) *CreateQosPolicyRequest {
	s.FlowIds = v
	return s
}

func (s *CreateQosPolicyRequest) SetInputRegionId(v string) *CreateQosPolicyRequest {
	s.InputRegionId = &v
	return s
}

func (s *CreateQosPolicyRequest) SetMaxIOBandWidth(v int64) *CreateQosPolicyRequest {
	s.MaxIOBandWidth = &v
	return s
}

func (s *CreateQosPolicyRequest) SetMaxIOps(v int64) *CreateQosPolicyRequest {
	s.MaxIOps = &v
	return s
}

func (s *CreateQosPolicyRequest) SetMaxMetaQps(v int64) *CreateQosPolicyRequest {
	s.MaxMetaQps = &v
	return s
}

func (s *CreateQosPolicyRequest) SetReqTags(v []*string) *CreateQosPolicyRequest {
	s.ReqTags = v
	return s
}

func (s *CreateQosPolicyRequest) SetZoneIds(v []*string) *CreateQosPolicyRequest {
	s.ZoneIds = v
	return s
}

func (s *CreateQosPolicyRequest) Validate() error {
	return dara.Validate(s)
}
