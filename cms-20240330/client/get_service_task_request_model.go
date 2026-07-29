// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *GetServiceTaskRequest
	GetType() *string
}

type GetServiceTaskRequest struct {
	// example:
	//
	// live_debug_log_probe
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetServiceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTaskRequest) GoString() string {
	return s.String()
}

func (s *GetServiceTaskRequest) GetType() *string {
	return s.Type
}

func (s *GetServiceTaskRequest) SetType(v string) *GetServiceTaskRequest {
	s.Type = &v
	return s
}

func (s *GetServiceTaskRequest) Validate() error {
	return dara.Validate(s)
}
