// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAvatarTrainingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateAvatarTrainingJobResponseBodyData) *UpdateAvatarTrainingJobResponseBody
	GetData() *UpdateAvatarTrainingJobResponseBodyData
	SetRequestId(v string) *UpdateAvatarTrainingJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAvatarTrainingJobResponseBody
	GetSuccess() *bool
}

type UpdateAvatarTrainingJobResponseBody struct {
	// The data returned.
	Data *UpdateAvatarTrainingJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAvatarTrainingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAvatarTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAvatarTrainingJobResponseBody) GetData() *UpdateAvatarTrainingJobResponseBodyData {
	return s.Data
}

func (s *UpdateAvatarTrainingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAvatarTrainingJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAvatarTrainingJobResponseBody) SetData(v *UpdateAvatarTrainingJobResponseBodyData) *UpdateAvatarTrainingJobResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAvatarTrainingJobResponseBody) SetRequestId(v string) *UpdateAvatarTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAvatarTrainingJobResponseBody) SetSuccess(v bool) *UpdateAvatarTrainingJobResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAvatarTrainingJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateAvatarTrainingJobResponseBodyData struct {
	// The ID of the digital human training job.
	//
	// example:
	//
	// ****d718e2ff4f018ccf419a7b71****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s UpdateAvatarTrainingJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateAvatarTrainingJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateAvatarTrainingJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *UpdateAvatarTrainingJobResponseBodyData) SetJobId(v string) *UpdateAvatarTrainingJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *UpdateAvatarTrainingJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
