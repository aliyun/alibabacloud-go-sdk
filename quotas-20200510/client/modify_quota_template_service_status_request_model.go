// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQuotaTemplateServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceStatus(v int32) *ModifyQuotaTemplateServiceStatusRequest
	GetServiceStatus() *int32
}

type ModifyQuotaTemplateServiceStatusRequest struct {
	// The status of the quota template. Valid values:
	//
	// 	- \\-1: The quota template is disabled.
	//
	// 	- 1: The quota template is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s ModifyQuotaTemplateServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyQuotaTemplateServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyQuotaTemplateServiceStatusRequest) GetServiceStatus() *int32 {
	return s.ServiceStatus
}

func (s *ModifyQuotaTemplateServiceStatusRequest) SetServiceStatus(v int32) *ModifyQuotaTemplateServiceStatusRequest {
	s.ServiceStatus = &v
	return s
}

func (s *ModifyQuotaTemplateServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
