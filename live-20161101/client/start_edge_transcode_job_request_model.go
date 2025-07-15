// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEdgeTranscodeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *StartEdgeTranscodeJobRequest
	GetClusterId() *string
	SetJobId(v string) *StartEdgeTranscodeJobRequest
	GetJobId() *string
	SetOwnerId(v int64) *StartEdgeTranscodeJobRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StartEdgeTranscodeJobRequest
	GetRegionId() *string
}

type StartEdgeTranscodeJobRequest struct {
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

func (s StartEdgeTranscodeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StartEdgeTranscodeJobRequest) GoString() string {
	return s.String()
}

func (s *StartEdgeTranscodeJobRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *StartEdgeTranscodeJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *StartEdgeTranscodeJobRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartEdgeTranscodeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartEdgeTranscodeJobRequest) SetClusterId(v string) *StartEdgeTranscodeJobRequest {
	s.ClusterId = &v
	return s
}

func (s *StartEdgeTranscodeJobRequest) SetJobId(v string) *StartEdgeTranscodeJobRequest {
	s.JobId = &v
	return s
}

func (s *StartEdgeTranscodeJobRequest) SetOwnerId(v int64) *StartEdgeTranscodeJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StartEdgeTranscodeJobRequest) SetRegionId(v string) *StartEdgeTranscodeJobRequest {
	s.RegionId = &v
	return s
}

func (s *StartEdgeTranscodeJobRequest) Validate() error {
	return dara.Validate(s)
}
