// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartWuyingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableRestartRetry(v bool) *RestartWuyingServerRequest
	GetEnableRestartRetry() *bool
	SetProductType(v string) *RestartWuyingServerRequest
	GetProductType() *string
	SetWuyingServerIdList(v []*string) *RestartWuyingServerRequest
	GetWuyingServerIdList() []*string
}

type RestartWuyingServerRequest struct {
	// Specifies whether to enable the restart retry feature.
	EnableRestartRetry *bool `json:"EnableRestartRetry,omitempty" xml:"EnableRestartRetry,omitempty"`
	// The product type.
	//
	// example:
	//
	// WuyingServer
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The list of workstation IDs.
	WuyingServerIdList []*string `json:"WuyingServerIdList,omitempty" xml:"WuyingServerIdList,omitempty" type:"Repeated"`
}

func (s RestartWuyingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartWuyingServerRequest) GoString() string {
	return s.String()
}

func (s *RestartWuyingServerRequest) GetEnableRestartRetry() *bool {
	return s.EnableRestartRetry
}

func (s *RestartWuyingServerRequest) GetProductType() *string {
	return s.ProductType
}

func (s *RestartWuyingServerRequest) GetWuyingServerIdList() []*string {
	return s.WuyingServerIdList
}

func (s *RestartWuyingServerRequest) SetEnableRestartRetry(v bool) *RestartWuyingServerRequest {
	s.EnableRestartRetry = &v
	return s
}

func (s *RestartWuyingServerRequest) SetProductType(v string) *RestartWuyingServerRequest {
	s.ProductType = &v
	return s
}

func (s *RestartWuyingServerRequest) SetWuyingServerIdList(v []*string) *RestartWuyingServerRequest {
	s.WuyingServerIdList = v
	return s
}

func (s *RestartWuyingServerRequest) Validate() error {
	return dara.Validate(s)
}
