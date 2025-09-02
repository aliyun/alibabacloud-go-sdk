// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceConsumeTimeRankRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizdate(v string) *GetInstanceConsumeTimeRankRequest
	GetBizdate() *string
	SetProjectId(v int64) *GetInstanceConsumeTimeRankRequest
	GetProjectId() *int64
}

type GetInstanceConsumeTimeRankRequest struct {
	// The data timestamp, accurate to the day. Specify the time in the ISO 8601 standard in the yyyy-MM-dd\\"T\\"HH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-21T00:00:00+0800
	Bizdate *string `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetInstanceConsumeTimeRankRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceConsumeTimeRankRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceConsumeTimeRankRequest) GetBizdate() *string {
	return s.Bizdate
}

func (s *GetInstanceConsumeTimeRankRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetInstanceConsumeTimeRankRequest) SetBizdate(v string) *GetInstanceConsumeTimeRankRequest {
	s.Bizdate = &v
	return s
}

func (s *GetInstanceConsumeTimeRankRequest) SetProjectId(v int64) *GetInstanceConsumeTimeRankRequest {
	s.ProjectId = &v
	return s
}

func (s *GetInstanceConsumeTimeRankRequest) Validate() error {
	return dara.Validate(s)
}
