// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeJobRequest
	GetAppId() *string
	SetJobId(v string) *DescribeJobRequest
	GetJobId() *string
}

type DescribeJobRequest struct {
	// The ID of the job template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// event-b798157b-40a2-4388-b578-71fb897103**-**
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeJobRequest) SetAppId(v string) *DescribeJobRequest {
	s.AppId = &v
	return s
}

func (s *DescribeJobRequest) SetJobId(v string) *DescribeJobRequest {
	s.JobId = &v
	return s
}

func (s *DescribeJobRequest) Validate() error {
	return dara.Validate(s)
}
