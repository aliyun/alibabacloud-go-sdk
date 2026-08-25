// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetServiceStatusResponseBody
	GetRequestId() *string
	SetServiceStatus(v *GetServiceStatusResponseBodyServiceStatus) *GetServiceStatusResponseBody
	GetServiceStatus() *GetServiceStatusResponseBodyServiceStatus
}

type GetServiceStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ADADC31D-90EE-5459-99B0-D83DF07769A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status information of CloudSSO.
	ServiceStatus *GetServiceStatusResponseBodyServiceStatus `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty" type:"Struct"`
}

func (s GetServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceStatusResponseBody) GetServiceStatus() *GetServiceStatusResponseBodyServiceStatus {
	return s.ServiceStatus
}

func (s *GetServiceStatusResponseBody) SetRequestId(v string) *GetServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceStatusResponseBody) SetServiceStatus(v *GetServiceStatusResponseBodyServiceStatus) *GetServiceStatusResponseBody {
	s.ServiceStatus = v
	return s
}

func (s *GetServiceStatusResponseBody) Validate() error {
	if s.ServiceStatus != nil {
		if err := s.ServiceStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceStatusResponseBodyServiceStatus struct {
	// The ID of your Alibaba Cloud account.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Indicates whether the prerequisites for enabling CloudSSO are met. Valid values:
	//
	// - Success: The prerequisites are met.
	//
	// - Failed: The prerequisites are not met.
	//
	// > The value of this parameter is returned only if the value of `Status` is `Disabled`.
	//
	// example:
	//
	// Success
	PrerequisiteCheckResult *string `json:"PrerequisiteCheckResult,omitempty" xml:"PrerequisiteCheckResult,omitempty"`
	// The IDs of regions where directories are deployed.
	RegionsInUse []*string `json:"RegionsInUse,omitempty" xml:"RegionsInUse,omitempty" type:"Repeated"`
	// The status of CloudSSO. Valid values:
	//
	// - Enabled
	//
	// - Disabled
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetServiceStatusResponseBodyServiceStatus) String() string {
	return dara.Prettify(s)
}

func (s GetServiceStatusResponseBodyServiceStatus) GoString() string {
	return s.String()
}

func (s *GetServiceStatusResponseBodyServiceStatus) GetAccountId() *string {
	return s.AccountId
}

func (s *GetServiceStatusResponseBodyServiceStatus) GetPrerequisiteCheckResult() *string {
	return s.PrerequisiteCheckResult
}

func (s *GetServiceStatusResponseBodyServiceStatus) GetRegionsInUse() []*string {
	return s.RegionsInUse
}

func (s *GetServiceStatusResponseBodyServiceStatus) GetStatus() *string {
	return s.Status
}

func (s *GetServiceStatusResponseBodyServiceStatus) SetAccountId(v string) *GetServiceStatusResponseBodyServiceStatus {
	s.AccountId = &v
	return s
}

func (s *GetServiceStatusResponseBodyServiceStatus) SetPrerequisiteCheckResult(v string) *GetServiceStatusResponseBodyServiceStatus {
	s.PrerequisiteCheckResult = &v
	return s
}

func (s *GetServiceStatusResponseBodyServiceStatus) SetRegionsInUse(v []*string) *GetServiceStatusResponseBodyServiceStatus {
	s.RegionsInUse = v
	return s
}

func (s *GetServiceStatusResponseBodyServiceStatus) SetStatus(v string) *GetServiceStatusResponseBodyServiceStatus {
	s.Status = &v
	return s
}

func (s *GetServiceStatusResponseBodyServiceStatus) Validate() error {
	return dara.Validate(s)
}
