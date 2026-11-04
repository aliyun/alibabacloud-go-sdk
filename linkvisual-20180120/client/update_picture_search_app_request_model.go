// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePictureSearchAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceId(v string) *UpdatePictureSearchAppRequest
	GetAppInstanceId() *string
	SetDescription(v string) *UpdatePictureSearchAppRequest
	GetDescription() *string
	SetName(v string) *UpdatePictureSearchAppRequest
	GetName() *string
}

type UpdatePictureSearchAppRequest struct {
	// Instance ID of the image search application instance. You can call the [QueryPictureSearchApps](206608) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e429560124b44b685095df81a27****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The description of the image search application instance.
	//
	// The description can contain Chinese characters, uppercase and lowercase letters, digits, underscores (_), and hyphens (-). It can be up to 1,024 characters in length, and each Chinese character counts as two characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 以图搜图Demo2
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the image search application instance.
	//
	// The name can contain Chinese characters, uppercase and lowercase letters, digits, underscores (_), and hyphens (-). It can be up to 64 characters in length, and each Chinese character counts as two characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 新以图搜图
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdatePictureSearchAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePictureSearchAppRequest) GoString() string {
	return s.String()
}

func (s *UpdatePictureSearchAppRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *UpdatePictureSearchAppRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePictureSearchAppRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePictureSearchAppRequest) SetAppInstanceId(v string) *UpdatePictureSearchAppRequest {
	s.AppInstanceId = &v
	return s
}

func (s *UpdatePictureSearchAppRequest) SetDescription(v string) *UpdatePictureSearchAppRequest {
	s.Description = &v
	return s
}

func (s *UpdatePictureSearchAppRequest) SetName(v string) *UpdatePictureSearchAppRequest {
	s.Name = &v
	return s
}

func (s *UpdatePictureSearchAppRequest) Validate() error {
	return dara.Validate(s)
}
