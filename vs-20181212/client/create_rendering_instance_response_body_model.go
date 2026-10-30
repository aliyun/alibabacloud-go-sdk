// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRenderingInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceId(v string) *CreateRenderingInstanceResponseBody
	GetRenderingInstanceId() *string
	SetRenderingInstanceIds(v []*string) *CreateRenderingInstanceResponseBody
	GetRenderingInstanceIds() []*string
	SetRequestId(v string) *CreateRenderingInstanceResponseBody
	GetRequestId() *string
}

type CreateRenderingInstanceResponseBody struct {
	// The instance ID of the cloud application service.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The instance ID list of the cloud application service.
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRenderingInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRenderingInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRenderingInstanceResponseBody) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *CreateRenderingInstanceResponseBody) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *CreateRenderingInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRenderingInstanceResponseBody) SetRenderingInstanceId(v string) *CreateRenderingInstanceResponseBody {
	s.RenderingInstanceId = &v
	return s
}

func (s *CreateRenderingInstanceResponseBody) SetRenderingInstanceIds(v []*string) *CreateRenderingInstanceResponseBody {
	s.RenderingInstanceIds = v
	return s
}

func (s *CreateRenderingInstanceResponseBody) SetRequestId(v string) *CreateRenderingInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRenderingInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
