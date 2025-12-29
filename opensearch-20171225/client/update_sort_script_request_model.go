// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSortScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateSortScriptRequest
	GetDescription() *string
}

type UpdateSortScriptRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateSortScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSortScriptRequest) GoString() string {
	return s.String()
}

func (s *UpdateSortScriptRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSortScriptRequest) SetDescription(v string) *UpdateSortScriptRequest {
	s.Description = &v
	return s
}

func (s *UpdateSortScriptRequest) Validate() error {
	return dara.Validate(s)
}
