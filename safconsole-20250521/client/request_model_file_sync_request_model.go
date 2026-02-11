// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestModelFileSyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *RequestModelFileSyncRequest
	GetProjectId() *int64
}

type RequestModelFileSyncRequest struct {
	// Project ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RequestModelFileSyncRequest) String() string {
	return dara.Prettify(s)
}

func (s RequestModelFileSyncRequest) GoString() string {
	return s.String()
}

func (s *RequestModelFileSyncRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *RequestModelFileSyncRequest) SetProjectId(v int64) *RequestModelFileSyncRequest {
	s.ProjectId = &v
	return s
}

func (s *RequestModelFileSyncRequest) Validate() error {
	return dara.Validate(s)
}
