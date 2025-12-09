// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRenderingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceIds(v []*string) *RebootRenderingServerRequest
	GetRenderingInstanceIds() []*string
}

type RebootRenderingServerRequest struct {
	// This parameter is required.
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
}

func (s RebootRenderingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootRenderingServerRequest) GoString() string {
	return s.String()
}

func (s *RebootRenderingServerRequest) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *RebootRenderingServerRequest) SetRenderingInstanceIds(v []*string) *RebootRenderingServerRequest {
	s.RenderingInstanceIds = v
	return s
}

func (s *RebootRenderingServerRequest) Validate() error {
	return dara.Validate(s)
}
