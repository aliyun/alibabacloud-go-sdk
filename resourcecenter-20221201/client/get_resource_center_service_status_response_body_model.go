// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceCenterServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInitialStatus(v string) *GetResourceCenterServiceStatusResponseBody
	GetInitialStatus() *string
	SetRequestId(v string) *GetResourceCenterServiceStatusResponseBody
	GetRequestId() *string
	SetServiceStatus(v string) *GetResourceCenterServiceStatusResponseBody
	GetServiceStatus() *string
}

type GetResourceCenterServiceStatusResponseBody struct {
	// The initialization status of the service. Valid values:
	//
	// 	- Pending: The service is being initialized.
	//
	// 	- Finished: The service is initialized.
	//
	// example:
	//
	// Pending
	InitialStatus *string `json:"InitialStatus,omitempty" xml:"InitialStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AD5F848D-CCDC-5464-93E1-4BA50A4826DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the service. Valid values:
	//
	// 	- Enabled: The service is activated.
	//
	// 	- Disabled: The service is deactivated.
	//
	// example:
	//
	// Enabled
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetResourceCenterServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCenterServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceCenterServiceStatusResponseBody) GetInitialStatus() *string {
	return s.InitialStatus
}

func (s *GetResourceCenterServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceCenterServiceStatusResponseBody) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *GetResourceCenterServiceStatusResponseBody) SetInitialStatus(v string) *GetResourceCenterServiceStatusResponseBody {
	s.InitialStatus = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponseBody) SetRequestId(v string) *GetResourceCenterServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponseBody) SetServiceStatus(v string) *GetResourceCenterServiceStatusResponseBody {
	s.ServiceStatus = &v
	return s
}

func (s *GetResourceCenterServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
