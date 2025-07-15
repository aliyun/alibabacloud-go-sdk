// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteAttendedTransferRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *CompleteAttendedTransferRequest
	GetDeviceId() *string
	SetInstanceId(v string) *CompleteAttendedTransferRequest
	GetInstanceId() *string
	SetJobId(v string) *CompleteAttendedTransferRequest
	GetJobId() *string
	SetUserId(v string) *CompleteAttendedTransferRequest
	GetUserId() *string
}

type CompleteAttendedTransferRequest struct {
	// example:
	//
	// device
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// job-6538214103685****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CompleteAttendedTransferRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteAttendedTransferRequest) GoString() string {
	return s.String()
}

func (s *CompleteAttendedTransferRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *CompleteAttendedTransferRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CompleteAttendedTransferRequest) GetJobId() *string {
	return s.JobId
}

func (s *CompleteAttendedTransferRequest) GetUserId() *string {
	return s.UserId
}

func (s *CompleteAttendedTransferRequest) SetDeviceId(v string) *CompleteAttendedTransferRequest {
	s.DeviceId = &v
	return s
}

func (s *CompleteAttendedTransferRequest) SetInstanceId(v string) *CompleteAttendedTransferRequest {
	s.InstanceId = &v
	return s
}

func (s *CompleteAttendedTransferRequest) SetJobId(v string) *CompleteAttendedTransferRequest {
	s.JobId = &v
	return s
}

func (s *CompleteAttendedTransferRequest) SetUserId(v string) *CompleteAttendedTransferRequest {
	s.UserId = &v
	return s
}

func (s *CompleteAttendedTransferRequest) Validate() error {
	return dara.Validate(s)
}
