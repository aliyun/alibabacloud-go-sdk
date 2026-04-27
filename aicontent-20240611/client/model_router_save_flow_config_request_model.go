// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterSaveFlowConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelId(v int32) *ModelRouterSaveFlowConfigRequest
	GetModelId() *int32
	SetRpm(v int32) *ModelRouterSaveFlowConfigRequest
	GetRpm() *int32
	SetSmoothFlowEnabled(v bool) *ModelRouterSaveFlowConfigRequest
	GetSmoothFlowEnabled() *bool
	SetTpm(v int32) *ModelRouterSaveFlowConfigRequest
	GetTpm() *int32
}

type ModelRouterSaveFlowConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 607
	ModelId *int32 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 100
	Rpm *int32 `json:"rpm,omitempty" xml:"rpm,omitempty"`
	// example:
	//
	// true
	SmoothFlowEnabled *bool `json:"smoothFlowEnabled,omitempty" xml:"smoothFlowEnabled,omitempty"`
	// example:
	//
	// 10000
	Tpm *int32 `json:"tpm,omitempty" xml:"tpm,omitempty"`
}

func (s ModelRouterSaveFlowConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterSaveFlowConfigRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterSaveFlowConfigRequest) GetModelId() *int32 {
	return s.ModelId
}

func (s *ModelRouterSaveFlowConfigRequest) GetRpm() *int32 {
	return s.Rpm
}

func (s *ModelRouterSaveFlowConfigRequest) GetSmoothFlowEnabled() *bool {
	return s.SmoothFlowEnabled
}

func (s *ModelRouterSaveFlowConfigRequest) GetTpm() *int32 {
	return s.Tpm
}

func (s *ModelRouterSaveFlowConfigRequest) SetModelId(v int32) *ModelRouterSaveFlowConfigRequest {
	s.ModelId = &v
	return s
}

func (s *ModelRouterSaveFlowConfigRequest) SetRpm(v int32) *ModelRouterSaveFlowConfigRequest {
	s.Rpm = &v
	return s
}

func (s *ModelRouterSaveFlowConfigRequest) SetSmoothFlowEnabled(v bool) *ModelRouterSaveFlowConfigRequest {
	s.SmoothFlowEnabled = &v
	return s
}

func (s *ModelRouterSaveFlowConfigRequest) SetTpm(v int32) *ModelRouterSaveFlowConfigRequest {
	s.Tpm = &v
	return s
}

func (s *ModelRouterSaveFlowConfigRequest) Validate() error {
	return dara.Validate(s)
}
