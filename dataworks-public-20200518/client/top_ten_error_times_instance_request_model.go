// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTopTenErrorTimesInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *TopTenErrorTimesInstanceRequest
	GetProjectId() *int64
}

type TopTenErrorTimesInstanceRequest struct {
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9527
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s TopTenErrorTimesInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s TopTenErrorTimesInstanceRequest) GoString() string {
	return s.String()
}

func (s *TopTenErrorTimesInstanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TopTenErrorTimesInstanceRequest) SetProjectId(v int64) *TopTenErrorTimesInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *TopTenErrorTimesInstanceRequest) Validate() error {
	return dara.Validate(s)
}
