// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegratedServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceCode(v string) *GetIntegratedServiceStatusRequest
	GetServiceCode() *string
}

type GetIntegratedServiceStatusRequest struct {
	// The identity of the cloud service that is integrated with Cloud Config. Valid values:
	//
	// 	- eventbridge: EventBridge
	//
	// 	- cms: CloudMonitor
	//
	// 	- bpstudio: Cloud Architect Design Tools (CADT)
	//
	// This parameter is required.
	//
	// example:
	//
	// cadt
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetIntegratedServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIntegratedServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetIntegratedServiceStatusRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetIntegratedServiceStatusRequest) SetServiceCode(v string) *GetIntegratedServiceStatusRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetIntegratedServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
