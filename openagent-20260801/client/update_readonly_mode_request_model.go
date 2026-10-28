// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReadonlyModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetValue(v string) *UpdateReadonlyModeRequest
	GetValue() *string
}

type UpdateReadonlyModeRequest struct {
	// Required. Pass the string true to enable read-only mode, or false to disable it. The value is case-insensitive.
	//
	// example:
	//
	// true
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateReadonlyModeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateReadonlyModeRequest) GoString() string {
	return s.String()
}

func (s *UpdateReadonlyModeRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateReadonlyModeRequest) SetValue(v string) *UpdateReadonlyModeRequest {
	s.Value = &v
	return s
}

func (s *UpdateReadonlyModeRequest) Validate() error {
	return dara.Validate(s)
}
