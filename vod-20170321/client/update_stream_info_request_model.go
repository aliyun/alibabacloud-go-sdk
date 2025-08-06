// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStreamInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpdateStreamInfoRequest
	GetJobId() *string
	SetMediaId(v string) *UpdateStreamInfoRequest
	GetMediaId() *string
}

type UpdateStreamInfoRequest struct {
	// This parameter is required.
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s UpdateStreamInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateStreamInfoRequest) GetJobId() *string {
	return s.JobId
}

func (s *UpdateStreamInfoRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateStreamInfoRequest) SetJobId(v string) *UpdateStreamInfoRequest {
	s.JobId = &v
	return s
}

func (s *UpdateStreamInfoRequest) SetMediaId(v string) *UpdateStreamInfoRequest {
	s.MediaId = &v
	return s
}

func (s *UpdateStreamInfoRequest) Validate() error {
	return dara.Validate(s)
}
