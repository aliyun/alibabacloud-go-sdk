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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only printable ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329241C
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the AI application.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
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
