// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateMemoryRequest
	GetDescription() *string
}

type CreateMemoryRequest struct {
	// The description of the long-term memory. The description is empty by default. The value must be 1 to 50 characters in length and can contain characters in the letter category of Unicode, which includes letters, Chinese characters, and digits. The value can also contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// Long-term memory of user A for my LLM application $APP_ID
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRequest) GoString() string {
	return s.String()
}

func (s *CreateMemoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMemoryRequest) SetDescription(v string) *CreateMemoryRequest {
	s.Description = &v
	return s
}

func (s *CreateMemoryRequest) Validate() error {
	return dara.Validate(s)
}
