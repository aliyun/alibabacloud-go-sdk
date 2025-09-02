// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTopTenElapsedTimeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *TopTenElapsedTimeInstanceRequest
	GetProjectId() *int64
}

type TopTenElapsedTimeInstanceRequest struct {
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Workspace page to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s TopTenElapsedTimeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s TopTenElapsedTimeInstanceRequest) GoString() string {
	return s.String()
}

func (s *TopTenElapsedTimeInstanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TopTenElapsedTimeInstanceRequest) SetProjectId(v int64) *TopTenElapsedTimeInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *TopTenElapsedTimeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
