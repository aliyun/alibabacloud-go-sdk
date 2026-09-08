// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAttendedTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *CancelAttendedTransferRequest
	GetDeviceId() *string
	SetInstanceId(v string) *CancelAttendedTransferRequest
	GetInstanceId() *string
	SetJobId(v string) *CancelAttendedTransferRequest
	GetJobId() *string
	SetUserId(v string) *CancelAttendedTransferRequest
	GetUserId() *string
}

type CancelAttendedTransferRequest struct {
	// The Device ID. This parameter is meaningless and can be filled with any value.
	//
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The call ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The agent ID.
	//
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CancelAttendedTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAttendedTransferRequest) GoString() string {
	return s.String()
}

func (s *CancelAttendedTransferRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *CancelAttendedTransferRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelAttendedTransferRequest) GetJobId() *string {
	return s.JobId
}

func (s *CancelAttendedTransferRequest) GetUserId() *string {
	return s.UserId
}

func (s *CancelAttendedTransferRequest) SetDeviceId(v string) *CancelAttendedTransferRequest {
	s.DeviceId = &v
	return s
}

func (s *CancelAttendedTransferRequest) SetInstanceId(v string) *CancelAttendedTransferRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelAttendedTransferRequest) SetJobId(v string) *CancelAttendedTransferRequest {
	s.JobId = &v
	return s
}

func (s *CancelAttendedTransferRequest) SetUserId(v string) *CancelAttendedTransferRequest {
	s.UserId = &v
	return s
}

func (s *CancelAttendedTransferRequest) Validate() error {
	return dara.Validate(s)
}
