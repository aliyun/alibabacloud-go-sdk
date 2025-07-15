// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateEdgeTranscodeJobRequest
	GetClusterId() *string
	SetJobId(v string) *UpdateEdgeTranscodeJobRequest
	GetJobId() *string
	SetName(v string) *UpdateEdgeTranscodeJobRequest
	GetName() *string
	SetOwnerId(v int64) *UpdateEdgeTranscodeJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateEdgeTranscodeJobRequest
	GetRegionId() *string
	SetStreamInput(v string) *UpdateEdgeTranscodeJobRequest
	GetStreamInput() *string
	SetStreamOutput(v string) *UpdateEdgeTranscodeJobRequest
	GetStreamOutput() *string
	SetTemplateId(v string) *UpdateEdgeTranscodeJobRequest
	GetTemplateId() *string
}

type UpdateEdgeTranscodeJobRequest struct {
	// The ID of the data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******3b-4d18-395c-8106-ff21a6******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the edge transcoding task.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The task name.
	//
	// example:
	//
	// task1
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The URL of the input stream.
	//
	// example:
	//
	// rtmp://mydomain/app/stream1
	StreamInput *string `json:"StreamInput,omitempty" xml:"StreamInput,omitempty"`
	// The URL of the output stream.
	//
	// example:
	//
	// rtmp://testdomain/app/stream2
	StreamOutput *string `json:"StreamOutput,omitempty" xml:"StreamOutput,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateEdgeTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateEdgeTranscodeJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateEdgeTranscodeJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdateEdgeTranscodeJobRequest) GetName() *string {
	return s.Name
}

func (s *UpdateEdgeTranscodeJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateEdgeTranscodeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateEdgeTranscodeJobRequest) GetStreamInput() *string {
	return s.StreamInput
}

func (s *UpdateEdgeTranscodeJobRequest) GetStreamOutput() *string {
	return s.StreamOutput
}

func (s *UpdateEdgeTranscodeJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateEdgeTranscodeJobRequest) SetClusterId(v string) *UpdateEdgeTranscodeJobRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateEdgeTranscodeJobRequest) SetJobId(v string) *UpdateEdgeTranscodeJobRequest {
	s.JobId = &v
	return s
}

func (s *UpdateEdgeTranscodeJobRequest) SetName(v string) *UpdateEdgeTranscodeJobRequest {
	s.Name = &v
	return s
}

func (s *UpdateEdgeTranscodeJobRequest) SetOwnerId(v int64) *UpdateEdgeTranscodeJobRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateEdgeTranscodeJobRequest) SetRegionId(v string) *UpdateEdgeTranscodeJobRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateEdgeTranscodeJobRequest) SetStreamInput(v string) *UpdateEdgeTranscodeJobRequest {
	s.StreamInput = &v
	return s
}

func (s *UpdateEdgeTranscodeJobRequest) SetStreamOutput(v string) *UpdateEdgeTranscodeJobRequest {
	s.StreamOutput = &v
	return s
}

func (s *UpdateEdgeTranscodeJobRequest) SetTemplateId(v string) *UpdateEdgeTranscodeJobRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateEdgeTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}
