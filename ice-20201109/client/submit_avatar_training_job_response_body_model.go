// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAvatarTrainingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitAvatarTrainingJobResponseBodyData) *SubmitAvatarTrainingJobResponseBody
	GetData() *SubmitAvatarTrainingJobResponseBodyData
	SetRequestId(v string) *SubmitAvatarTrainingJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitAvatarTrainingJobResponseBody
	GetSuccess() *bool
}

type SubmitAvatarTrainingJobResponseBody struct {
	// The data returned.
	Data *SubmitAvatarTrainingJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// *****ACB-44F2-5F2D-88D7-1283E70*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAvatarTrainingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAvatarTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAvatarTrainingJobResponseBody) GetData() *SubmitAvatarTrainingJobResponseBodyData {
	return s.Data
}

func (s *SubmitAvatarTrainingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAvatarTrainingJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitAvatarTrainingJobResponseBody) SetData(v *SubmitAvatarTrainingJobResponseBodyData) *SubmitAvatarTrainingJobResponseBody {
	s.Data = v
	return s
}

func (s *SubmitAvatarTrainingJobResponseBody) SetRequestId(v string) *SubmitAvatarTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAvatarTrainingJobResponseBody) SetSuccess(v bool) *SubmitAvatarTrainingJobResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitAvatarTrainingJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitAvatarTrainingJobResponseBodyData struct {
	// The ID of the digital human training job.
	//
	// example:
	//
	// ****29faef8144638ba42eb8e037****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitAvatarTrainingJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitAvatarTrainingJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitAvatarTrainingJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAvatarTrainingJobResponseBodyData) SetJobId(v string) *SubmitAvatarTrainingJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *SubmitAvatarTrainingJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
