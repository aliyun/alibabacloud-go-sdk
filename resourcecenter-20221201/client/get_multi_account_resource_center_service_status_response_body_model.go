// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountResourceCenterServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInitialStatus(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody
	GetInitialStatus() *string
	SetRequestId(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody
	GetRequestId() *string
	SetServiceStatus(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody
	GetServiceStatus() *string
}

type GetMultiAccountResourceCenterServiceStatusResponseBody struct {
	// The initialization status of the feature. Valid values:
	//
	// 	- Pending: The feature is being initialized.
	//
	// 	- Finished: The feature is initialized.
	//
	// example:
	//
	// Pending
	InitialStatus *string `json:"InitialStatus,omitempty" xml:"InitialStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 81671397-1425-51F1-A144-4799E01BEBFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the feature. Valid values:
	//
	// 	- Enabled: The feature is enabled.
	//
	// 	- Disabled: The feature is disabled.
	//
	// example:
	//
	// Enabled
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetMultiAccountResourceCenterServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCenterServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) GetInitialStatus() *string {
	return s.InitialStatus
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetInitialStatus(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.InitialStatus = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetRequestId(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) SetServiceStatus(v string) *GetMultiAccountResourceCenterServiceStatusResponseBody {
	s.ServiceStatus = &v
	return s
}

func (s *GetMultiAccountResourceCenterServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
