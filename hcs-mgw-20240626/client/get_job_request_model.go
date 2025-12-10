// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetByVersion(v string) *GetJobRequest
	GetByVersion() *string
}

type GetJobRequest struct {
	// Specifies whether to obtain the details of the migration task by using the task ID.
	//
	// example:
	//
	// false
	ByVersion *string `json:"byVersion,omitempty" xml:"byVersion,omitempty"`
}

func (s GetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) GetByVersion() *string {
	return s.ByVersion
}

func (s *GetJobRequest) SetByVersion(v string) *GetJobRequest {
	s.ByVersion = &v
	return s
}

func (s *GetJobRequest) Validate() error {
	return dara.Validate(s)
}
