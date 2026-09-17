// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRenderingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrecheck(v bool) *RebootRenderingServerRequest
	GetPrecheck() *bool
	SetRenderingInstanceIds(v []*string) *RebootRenderingServerRequest
	GetRenderingInstanceIds() []*string
}

type RebootRenderingServerRequest struct {
	// Specifies whether to perform only an admission check without actually restarting the hosts. Default value: false.
	//
	// example:
	//
	// true
	Precheck *bool `json:"Precheck,omitempty" xml:"Precheck,omitempty"`
	// The list of cloud application service instance IDs.
	//
	// This parameter is required.
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
}

func (s RebootRenderingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootRenderingServerRequest) GoString() string {
	return s.String()
}

func (s *RebootRenderingServerRequest) GetPrecheck() *bool {
	return s.Precheck
}

func (s *RebootRenderingServerRequest) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *RebootRenderingServerRequest) SetPrecheck(v bool) *RebootRenderingServerRequest {
	s.Precheck = &v
	return s
}

func (s *RebootRenderingServerRequest) SetRenderingInstanceIds(v []*string) *RebootRenderingServerRequest {
	s.RenderingInstanceIds = v
	return s
}

func (s *RebootRenderingServerRequest) Validate() error {
	return dara.Validate(s)
}
