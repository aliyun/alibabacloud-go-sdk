// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *AnswerCallRequest
	GetDeviceId() *string
	SetInstanceId(v string) *AnswerCallRequest
	GetInstanceId() *string
	SetJobId(v string) *AnswerCallRequest
	GetJobId() *string
	SetUserId(v string) *AnswerCallRequest
	GetUserId() *string
}

type AnswerCallRequest struct {
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
	// job-65382141036853491
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AnswerCallRequest) String() string {
	return dara.Prettify(s)
}

func (s AnswerCallRequest) GoString() string {
	return s.String()
}

func (s *AnswerCallRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *AnswerCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AnswerCallRequest) GetJobId() *string {
	return s.JobId
}

func (s *AnswerCallRequest) GetUserId() *string {
	return s.UserId
}

func (s *AnswerCallRequest) SetDeviceId(v string) *AnswerCallRequest {
	s.DeviceId = &v
	return s
}

func (s *AnswerCallRequest) SetInstanceId(v string) *AnswerCallRequest {
	s.InstanceId = &v
	return s
}

func (s *AnswerCallRequest) SetJobId(v string) *AnswerCallRequest {
	s.JobId = &v
	return s
}

func (s *AnswerCallRequest) SetUserId(v string) *AnswerCallRequest {
	s.UserId = &v
	return s
}

func (s *AnswerCallRequest) Validate() error {
	return dara.Validate(s)
}
