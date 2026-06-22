// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotQueryStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwner(v string) *KopilotQueryStatusRequest
	GetOwner() *string
	SetRegionId(v string) *KopilotQueryStatusRequest
	GetRegionId() *string
}

type KopilotQueryStatusRequest struct {
	// This parameter is required.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s KopilotQueryStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s KopilotQueryStatusRequest) GoString() string {
	return s.String()
}

func (s *KopilotQueryStatusRequest) GetOwner() *string {
	return s.Owner
}

func (s *KopilotQueryStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KopilotQueryStatusRequest) SetOwner(v string) *KopilotQueryStatusRequest {
	s.Owner = &v
	return s
}

func (s *KopilotQueryStatusRequest) SetRegionId(v string) *KopilotQueryStatusRequest {
	s.RegionId = &v
	return s
}

func (s *KopilotQueryStatusRequest) Validate() error {
	return dara.Validate(s)
}
