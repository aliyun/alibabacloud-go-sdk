// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationBaseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateApplicationBaseInfoRequest
	GetAppId() *string
	SetAppName(v string) *UpdateApplicationBaseInfoRequest
	GetAppName() *string
	SetDesc(v string) *UpdateApplicationBaseInfoRequest
	GetDesc() *string
	SetOwner(v string) *UpdateApplicationBaseInfoRequest
	GetOwner() *string
}

type UpdateApplicationBaseInfoRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// c627c157-560d-43ff-****-************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application. The name must start with a letter, and can contain letters, digits, underscores (_), and hyphens (-). The name can be up to 36 characters in length.
	//
	// example:
	//
	// hello-edas
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the application. The description can be up to 256 characters in length.
	//
	// example:
	//
	// Test application
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The owner of the application. The value can be up to 127 characters in length.
	//
	// example:
	//
	// test@aliyun_xxx.com
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s UpdateApplicationBaseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationBaseInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateApplicationBaseInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateApplicationBaseInfoRequest) GetDesc() *string {
	return s.Desc
}

func (s *UpdateApplicationBaseInfoRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateApplicationBaseInfoRequest) SetAppId(v string) *UpdateApplicationBaseInfoRequest {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationBaseInfoRequest) SetAppName(v string) *UpdateApplicationBaseInfoRequest {
	s.AppName = &v
	return s
}

func (s *UpdateApplicationBaseInfoRequest) SetDesc(v string) *UpdateApplicationBaseInfoRequest {
	s.Desc = &v
	return s
}

func (s *UpdateApplicationBaseInfoRequest) SetOwner(v string) *UpdateApplicationBaseInfoRequest {
	s.Owner = &v
	return s
}

func (s *UpdateApplicationBaseInfoRequest) Validate() error {
	return dara.Validate(s)
}
