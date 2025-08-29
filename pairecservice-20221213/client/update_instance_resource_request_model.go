// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateInstanceResourceRequest
	GetConfig() *string
	SetUri(v string) *UpdateInstanceResourceRequest
	GetUri() *string
}

type UpdateInstanceResourceRequest struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// bucket-test-123
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s UpdateInstanceResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceResourceRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateInstanceResourceRequest) GetUri() *string {
	return s.Uri
}

func (s *UpdateInstanceResourceRequest) SetConfig(v string) *UpdateInstanceResourceRequest {
	s.Config = &v
	return s
}

func (s *UpdateInstanceResourceRequest) SetUri(v string) *UpdateInstanceResourceRequest {
	s.Uri = &v
	return s
}

func (s *UpdateInstanceResourceRequest) Validate() error {
	return dara.Validate(s)
}
