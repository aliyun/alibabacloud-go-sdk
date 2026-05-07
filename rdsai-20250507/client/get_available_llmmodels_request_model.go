// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAvailableLLMModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetAvailableLLMModelsRequest
	GetClientToken() *string
	SetInstanceName(v string) *GetAvailableLLMModelsRequest
	GetInstanceName() *string
	SetRegionId(v string) *GetAvailableLLMModelsRequest
	GetRegionId() *string
}

type GetAvailableLLMModelsRequest struct {
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAvailableLLMModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAvailableLLMModelsRequest) GoString() string {
	return s.String()
}

func (s *GetAvailableLLMModelsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetAvailableLLMModelsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetAvailableLLMModelsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAvailableLLMModelsRequest) SetClientToken(v string) *GetAvailableLLMModelsRequest {
	s.ClientToken = &v
	return s
}

func (s *GetAvailableLLMModelsRequest) SetInstanceName(v string) *GetAvailableLLMModelsRequest {
	s.InstanceName = &v
	return s
}

func (s *GetAvailableLLMModelsRequest) SetRegionId(v string) *GetAvailableLLMModelsRequest {
	s.RegionId = &v
	return s
}

func (s *GetAvailableLLMModelsRequest) Validate() error {
	return dara.Validate(s)
}
