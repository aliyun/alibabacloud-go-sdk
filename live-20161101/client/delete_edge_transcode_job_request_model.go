// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEdgeTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteEdgeTranscodeJobRequest
	GetClusterId() *string
	SetJobId(v string) *DeleteEdgeTranscodeJobRequest
	GetJobId() *string
	SetOwnerId(v int64) *DeleteEdgeTranscodeJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteEdgeTranscodeJobRequest
	GetRegionId() *string
}

type DeleteEdgeTranscodeJobRequest struct {
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

func (s DeleteEdgeTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEdgeTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteEdgeTranscodeJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteEdgeTranscodeJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DeleteEdgeTranscodeJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteEdgeTranscodeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEdgeTranscodeJobRequest) SetClusterId(v string) *DeleteEdgeTranscodeJobRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteEdgeTranscodeJobRequest) SetJobId(v string) *DeleteEdgeTranscodeJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteEdgeTranscodeJobRequest) SetOwnerId(v int64) *DeleteEdgeTranscodeJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteEdgeTranscodeJobRequest) SetRegionId(v string) *DeleteEdgeTranscodeJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEdgeTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}
