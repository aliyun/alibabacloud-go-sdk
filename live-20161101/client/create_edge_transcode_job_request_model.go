// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateEdgeTranscodeJobRequest
	GetClusterId() *string
	SetName(v string) *CreateEdgeTranscodeJobRequest
	GetName() *string
	SetOwnerId(v int64) *CreateEdgeTranscodeJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateEdgeTranscodeJobRequest
	GetRegionId() *string
	SetStreamInput(v string) *CreateEdgeTranscodeJobRequest
	GetStreamInput() *string
	SetStreamOutput(v string) *CreateEdgeTranscodeJobRequest
	GetStreamOutput() *string
	SetTemplateId(v string) *CreateEdgeTranscodeJobRequest
	GetTemplateId() *string
}

type CreateEdgeTranscodeJobRequest struct {
	// The ID of the data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// ******3b-4d18-395c-8106-ff21a6******
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The task name. The name can contain letters, digits, hyphens (-), and underscores (_). The name must be 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// task1
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The URL of the input stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://mydomain/app/stream1
	StreamInput *string `json:"StreamInput,omitempty" xml:"StreamInput,omitempty"`
	// The URL of the output stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://testdomain/app/stream2
	StreamOutput *string `json:"StreamOutput,omitempty" xml:"StreamOutput,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateEdgeTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *CreateEdgeTranscodeJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateEdgeTranscodeJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateEdgeTranscodeJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateEdgeTranscodeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEdgeTranscodeJobRequest) GetStreamInput() *string {
	return s.StreamInput
}

func (s *CreateEdgeTranscodeJobRequest) GetStreamOutput() *string {
	return s.StreamOutput
}

func (s *CreateEdgeTranscodeJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateEdgeTranscodeJobRequest) SetClusterId(v string) *CreateEdgeTranscodeJobRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateEdgeTranscodeJobRequest) SetName(v string) *CreateEdgeTranscodeJobRequest {
	s.Name = &v
	return s
}

func (s *CreateEdgeTranscodeJobRequest) SetOwnerId(v int64) *CreateEdgeTranscodeJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateEdgeTranscodeJobRequest) SetRegionId(v string) *CreateEdgeTranscodeJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEdgeTranscodeJobRequest) SetStreamInput(v string) *CreateEdgeTranscodeJobRequest {
	s.StreamInput = &v
	return s
}

func (s *CreateEdgeTranscodeJobRequest) SetStreamOutput(v string) *CreateEdgeTranscodeJobRequest {
	s.StreamOutput = &v
	return s
}

func (s *CreateEdgeTranscodeJobRequest) SetTemplateId(v string) *CreateEdgeTranscodeJobRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateEdgeTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}
