// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPassword(v string) *ModifyRenderingInstanceAttributeRequest
	GetPassword() *string
	SetRenderingInstanceId(v string) *ModifyRenderingInstanceAttributeRequest
	GetRenderingInstanceId() *string
}

type ModifyRenderingInstanceAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Toehold2020
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s ModifyRenderingInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyRenderingInstanceAttributeRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyRenderingInstanceAttributeRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ModifyRenderingInstanceAttributeRequest) SetPassword(v string) *ModifyRenderingInstanceAttributeRequest {
	s.Password = &v
	return s
}

func (s *ModifyRenderingInstanceAttributeRequest) SetRenderingInstanceId(v string) *ModifyRenderingInstanceAttributeRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ModifyRenderingInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
