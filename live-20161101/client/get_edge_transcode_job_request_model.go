// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetEdgeTranscodeJobRequest
	GetClusterId() *string
	SetJobId(v string) *GetEdgeTranscodeJobRequest
	GetJobId() *string
	SetOwnerId(v int64) *GetEdgeTranscodeJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetEdgeTranscodeJobRequest
	GetRegionId() *string
}

type GetEdgeTranscodeJobRequest struct {
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

func (s GetEdgeTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *GetEdgeTranscodeJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetEdgeTranscodeJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetEdgeTranscodeJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetEdgeTranscodeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEdgeTranscodeJobRequest) SetClusterId(v string) *GetEdgeTranscodeJobRequest {
	s.ClusterId = &v
	return s
}

func (s *GetEdgeTranscodeJobRequest) SetJobId(v string) *GetEdgeTranscodeJobRequest {
	s.JobId = &v
	return s
}

func (s *GetEdgeTranscodeJobRequest) SetOwnerId(v int64) *GetEdgeTranscodeJobRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEdgeTranscodeJobRequest) SetRegionId(v string) *GetEdgeTranscodeJobRequest {
	s.RegionId = &v
	return s
}

func (s *GetEdgeTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}
