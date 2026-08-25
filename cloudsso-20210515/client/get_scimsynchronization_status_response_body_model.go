// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSCIMSynchronizationStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSCIMSynchronizationStatusResponseBody
	GetRequestId() *string
	SetSCIMSynchronizationStatus(v string) *GetSCIMSynchronizationStatusResponseBody
	GetSCIMSynchronizationStatus() *string
}

type GetSCIMSynchronizationStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7C086C2F-1C66-57B3-B14E-2C1DA70727CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of SCIM synchronization. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Enabled
	SCIMSynchronizationStatus *string `json:"SCIMSynchronizationStatus,omitempty" xml:"SCIMSynchronizationStatus,omitempty"`
}

func (s GetSCIMSynchronizationStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSCIMSynchronizationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetSCIMSynchronizationStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSCIMSynchronizationStatusResponseBody) GetSCIMSynchronizationStatus() *string {
	return s.SCIMSynchronizationStatus
}

func (s *GetSCIMSynchronizationStatusResponseBody) SetRequestId(v string) *GetSCIMSynchronizationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSCIMSynchronizationStatusResponseBody) SetSCIMSynchronizationStatus(v string) *GetSCIMSynchronizationStatusResponseBody {
	s.SCIMSynchronizationStatus = &v
	return s
}

func (s *GetSCIMSynchronizationStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
