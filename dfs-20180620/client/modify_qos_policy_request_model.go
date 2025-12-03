// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyQosPolicyRequest
	GetDescription() *string
	SetFederationId(v string) *ModifyQosPolicyRequest
	GetFederationId() *string
	SetFileSystemId(v string) *ModifyQosPolicyRequest
	GetFileSystemId() *string
	SetInputRegionId(v string) *ModifyQosPolicyRequest
	GetInputRegionId() *string
	SetMaxIOBandWidth(v int64) *ModifyQosPolicyRequest
	GetMaxIOBandWidth() *int64
	SetMaxIOps(v int64) *ModifyQosPolicyRequest
	GetMaxIOps() *int64
	SetMaxMetaQps(v int64) *ModifyQosPolicyRequest
	GetMaxMetaQps() *int64
	SetQosPolicyId(v string) *ModifyQosPolicyRequest
	GetQosPolicyId() *string
}

type ModifyQosPolicyRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FederationId *string `json:"FederationId,omitempty" xml:"FederationId,omitempty"`
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	InputRegionId  *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	MaxIOBandWidth *int64  `json:"MaxIOBandWidth,omitempty" xml:"MaxIOBandWidth,omitempty"`
	MaxIOps        *int64  `json:"MaxIOps,omitempty" xml:"MaxIOps,omitempty"`
	MaxMetaQps     *int64  `json:"MaxMetaQps,omitempty" xml:"MaxMetaQps,omitempty"`
	// This parameter is required.
	QosPolicyId *string `json:"QosPolicyId,omitempty" xml:"QosPolicyId,omitempty"`
}

func (s ModifyQosPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyQosPolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyQosPolicyRequest) GetFederationId() *string {
	return s.FederationId
}

func (s *ModifyQosPolicyRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyQosPolicyRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *ModifyQosPolicyRequest) GetMaxIOBandWidth() *int64 {
	return s.MaxIOBandWidth
}

func (s *ModifyQosPolicyRequest) GetMaxIOps() *int64 {
	return s.MaxIOps
}

func (s *ModifyQosPolicyRequest) GetMaxMetaQps() *int64 {
	return s.MaxMetaQps
}

func (s *ModifyQosPolicyRequest) GetQosPolicyId() *string {
	return s.QosPolicyId
}

func (s *ModifyQosPolicyRequest) SetDescription(v string) *ModifyQosPolicyRequest {
	s.Description = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetFederationId(v string) *ModifyQosPolicyRequest {
	s.FederationId = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetFileSystemId(v string) *ModifyQosPolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetInputRegionId(v string) *ModifyQosPolicyRequest {
	s.InputRegionId = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetMaxIOBandWidth(v int64) *ModifyQosPolicyRequest {
	s.MaxIOBandWidth = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetMaxIOps(v int64) *ModifyQosPolicyRequest {
	s.MaxIOps = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetMaxMetaQps(v int64) *ModifyQosPolicyRequest {
	s.MaxMetaQps = &v
	return s
}

func (s *ModifyQosPolicyRequest) SetQosPolicyId(v string) *ModifyQosPolicyRequest {
	s.QosPolicyId = &v
	return s
}

func (s *ModifyQosPolicyRequest) Validate() error {
	return dara.Validate(s)
}
