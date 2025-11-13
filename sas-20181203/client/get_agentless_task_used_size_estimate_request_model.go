// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentlessTaskUsedSizeEstimateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetSelectionType(v string) *GetAgentlessTaskUsedSizeEstimateRequest
	GetAssetSelectionType() *string
}

type GetAgentlessTaskUsedSizeEstimateRequest struct {
	// Asset selection identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// AGENTLESS_SCAN_ONCE_TASK_1720145******
	AssetSelectionType *string `json:"AssetSelectionType,omitempty" xml:"AssetSelectionType,omitempty"`
}

func (s GetAgentlessTaskUsedSizeEstimateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentlessTaskUsedSizeEstimateRequest) GoString() string {
	return s.String()
}

func (s *GetAgentlessTaskUsedSizeEstimateRequest) GetAssetSelectionType() *string {
	return s.AssetSelectionType
}

func (s *GetAgentlessTaskUsedSizeEstimateRequest) SetAssetSelectionType(v string) *GetAgentlessTaskUsedSizeEstimateRequest {
	s.AssetSelectionType = &v
	return s
}

func (s *GetAgentlessTaskUsedSizeEstimateRequest) Validate() error {
	return dara.Validate(s)
}
