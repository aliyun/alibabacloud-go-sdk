// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAvatarTrainingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateAvatarTrainingJobResponseBodyData) *CreateAvatarTrainingJobResponseBody
	GetData() *CreateAvatarTrainingJobResponseBodyData
	SetRequestId(v string) *CreateAvatarTrainingJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAvatarTrainingJobResponseBody
	GetSuccess() *bool
}

type CreateAvatarTrainingJobResponseBody struct {
	// The data returned.
	Data *CreateAvatarTrainingJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAvatarTrainingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAvatarTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAvatarTrainingJobResponseBody) GetData() *CreateAvatarTrainingJobResponseBodyData {
	return s.Data
}

func (s *CreateAvatarTrainingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAvatarTrainingJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAvatarTrainingJobResponseBody) SetData(v *CreateAvatarTrainingJobResponseBodyData) *CreateAvatarTrainingJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateAvatarTrainingJobResponseBody) SetRequestId(v string) *CreateAvatarTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAvatarTrainingJobResponseBody) SetSuccess(v bool) *CreateAvatarTrainingJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAvatarTrainingJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAvatarTrainingJobResponseBodyData struct {
	// The ID of the digital human training job.
	//
	// example:
	//
	// ****d718e2ff4f018ccf419a7b71****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s CreateAvatarTrainingJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAvatarTrainingJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAvatarTrainingJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *CreateAvatarTrainingJobResponseBodyData) SetJobId(v string) *CreateAvatarTrainingJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateAvatarTrainingJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
