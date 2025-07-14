// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeJobStatusRequest
	GetAppId() *string
	SetJobId(v string) *DescribeJobStatusRequest
	GetJobId() *string
}

type DescribeJobStatusRequest struct {
	// The ID of the job template.
	//
	// This parameter is required.
	//
	// example:
	//
	// e1a7a07-abcb-4652-a1d3-2d57f415****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// event-b798157b-40a2-4388-b578-71fb897103**-**
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeJobStatusRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobStatusRequest) SetAppId(v string) *DescribeJobStatusRequest {
	s.AppId = &v
	return s
}

func (s *DescribeJobStatusRequest) SetJobId(v string) *DescribeJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
