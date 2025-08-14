// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAudioProduceJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitAudioProduceJobResponseBody
	GetJobId() *string
	SetMediaId(v string) *SubmitAudioProduceJobResponseBody
	GetMediaId() *string
	SetRequestId(v string) *SubmitAudioProduceJobResponseBody
	GetRequestId() *string
	SetState(v string) *SubmitAudioProduceJobResponseBody
	GetState() *string
}

type SubmitAudioProduceJobResponseBody struct {
	// The job ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****2bcbfcfa30fccb36f72dca22****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The job state. Valid values:
	//
	// 	- Created
	//
	// 	- Executing
	//
	// 	- Finished
	//
	// 	- Failed
	//
	// example:
	//
	// Created
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SubmitAudioProduceJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAudioProduceJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAudioProduceJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAudioProduceJobResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *SubmitAudioProduceJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAudioProduceJobResponseBody) GetState() *string {
	return s.State
}

func (s *SubmitAudioProduceJobResponseBody) SetJobId(v string) *SubmitAudioProduceJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAudioProduceJobResponseBody) SetMediaId(v string) *SubmitAudioProduceJobResponseBody {
	s.MediaId = &v
	return s
}

func (s *SubmitAudioProduceJobResponseBody) SetRequestId(v string) *SubmitAudioProduceJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAudioProduceJobResponseBody) SetState(v string) *SubmitAudioProduceJobResponseBody {
	s.State = &v
	return s
}

func (s *SubmitAudioProduceJobResponseBody) Validate() error {
	return dara.Validate(s)
}
