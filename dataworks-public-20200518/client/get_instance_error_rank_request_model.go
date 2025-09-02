// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceErrorRankRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *GetInstanceErrorRankRequest
	GetProjectId() *int64
}

type GetInstanceErrorRankRequest struct {
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9527
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetInstanceErrorRankRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceErrorRankRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceErrorRankRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetInstanceErrorRankRequest) SetProjectId(v int64) *GetInstanceErrorRankRequest {
	s.ProjectId = &v
	return s
}

func (s *GetInstanceErrorRankRequest) Validate() error {
	return dara.Validate(s)
}
