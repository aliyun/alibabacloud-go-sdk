// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateMemoryRequest
	GetDescription() *string
}

type UpdateMemoryRequest struct {
	// The new description. The description must be 1 to 50 characters in length and can contain characters in the letter category of Unicode, which includes letters, Chinese characters, and digits. The description can also contain half-width colons (:), underscores (_), periods (.), or hyphens (-).
	//
	// example:
	//
	// 我的大模型应用$APP_ID关于B用户的长期记忆体
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMemoryRequest) SetDescription(v string) *UpdateMemoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateMemoryRequest) Validate() error {
	return dara.Validate(s)
}
