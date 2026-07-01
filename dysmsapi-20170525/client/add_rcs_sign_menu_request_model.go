// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRcsSignMenuRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMenuContent(v string) *AddRcsSignMenuRequest
	GetMenuContent() *string
	SetSignName(v string) *AddRcsSignMenuRequest
	GetSignName() *string
}

type AddRcsSignMenuRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	MenuContent *string `json:"MenuContent,omitempty" xml:"MenuContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
}

func (s AddRcsSignMenuRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRcsSignMenuRequest) GoString() string {
	return s.String()
}

func (s *AddRcsSignMenuRequest) GetMenuContent() *string {
	return s.MenuContent
}

func (s *AddRcsSignMenuRequest) GetSignName() *string {
	return s.SignName
}

func (s *AddRcsSignMenuRequest) SetMenuContent(v string) *AddRcsSignMenuRequest {
	s.MenuContent = &v
	return s
}

func (s *AddRcsSignMenuRequest) SetSignName(v string) *AddRcsSignMenuRequest {
	s.SignName = &v
	return s
}

func (s *AddRcsSignMenuRequest) Validate() error {
	return dara.Validate(s)
}
