// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopEdgeTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *StopEdgeTranscodeJobRequest
	GetClusterId() *string
	SetJobId(v string) *StopEdgeTranscodeJobRequest
	GetJobId() *string
	SetOwnerId(v int64) *StopEdgeTranscodeJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StopEdgeTranscodeJobRequest
	GetRegionId() *string
}

type StopEdgeTranscodeJobRequest struct {
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
	JobId    *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopEdgeTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopEdgeTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *StopEdgeTranscodeJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *StopEdgeTranscodeJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *StopEdgeTranscodeJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopEdgeTranscodeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopEdgeTranscodeJobRequest) SetClusterId(v string) *StopEdgeTranscodeJobRequest {
	s.ClusterId = &v
	return s
}

func (s *StopEdgeTranscodeJobRequest) SetJobId(v string) *StopEdgeTranscodeJobRequest {
	s.JobId = &v
	return s
}

func (s *StopEdgeTranscodeJobRequest) SetOwnerId(v int64) *StopEdgeTranscodeJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StopEdgeTranscodeJobRequest) SetRegionId(v string) *StopEdgeTranscodeJobRequest {
	s.RegionId = &v
	return s
}

func (s *StopEdgeTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}
