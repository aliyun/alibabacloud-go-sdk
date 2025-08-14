// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAvatarTrainingJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteAvatarTrainingJobResponseBodyData) *DeleteAvatarTrainingJobResponseBody
	GetData() *DeleteAvatarTrainingJobResponseBodyData
	SetRequestId(v string) *DeleteAvatarTrainingJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAvatarTrainingJobResponseBody
	GetSuccess() *bool
}

type DeleteAvatarTrainingJobResponseBody struct {
	// The data returned.
	Data *DeleteAvatarTrainingJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DeleteAvatarTrainingJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAvatarTrainingJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAvatarTrainingJobResponseBody) GetData() *DeleteAvatarTrainingJobResponseBodyData {
	return s.Data
}

func (s *DeleteAvatarTrainingJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAvatarTrainingJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAvatarTrainingJobResponseBody) SetData(v *DeleteAvatarTrainingJobResponseBodyData) *DeleteAvatarTrainingJobResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAvatarTrainingJobResponseBody) SetRequestId(v string) *DeleteAvatarTrainingJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAvatarTrainingJobResponseBody) SetSuccess(v bool) *DeleteAvatarTrainingJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAvatarTrainingJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteAvatarTrainingJobResponseBodyData struct {
	// The ID of the digital human training job.
	//
	// example:
	//
	// ****d718e2ff4f018ccf419a7b71****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteAvatarTrainingJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteAvatarTrainingJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteAvatarTrainingJobResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *DeleteAvatarTrainingJobResponseBodyData) SetJobId(v string) *DeleteAvatarTrainingJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *DeleteAvatarTrainingJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
