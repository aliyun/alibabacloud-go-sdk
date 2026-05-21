// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTempFileTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGmtExpiredTime(v string) *UpdateTempFileTaskRequest
	GetGmtExpiredTime() *string
}

type UpdateTempFileTaskRequest struct {
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtExpiredTime *string `json:"GmtExpiredTime,omitempty" xml:"GmtExpiredTime,omitempty"`
}

func (s UpdateTempFileTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTempFileTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateTempFileTaskRequest) GetGmtExpiredTime() *string {
	return s.GmtExpiredTime
}

func (s *UpdateTempFileTaskRequest) SetGmtExpiredTime(v string) *UpdateTempFileTaskRequest {
	s.GmtExpiredTime = &v
	return s
}

func (s *UpdateTempFileTaskRequest) Validate() error {
	return dara.Validate(s)
}
