// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExtensionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExtensionCode(v string) *GetExtensionRequest
	GetExtensionCode() *string
}

type GetExtensionRequest struct {
	// The unique code of the extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce4*********086da5
	ExtensionCode *string `json:"ExtensionCode,omitempty" xml:"ExtensionCode,omitempty"`
}

func (s GetExtensionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExtensionRequest) GoString() string {
	return s.String()
}

func (s *GetExtensionRequest) GetExtensionCode() *string {
	return s.ExtensionCode
}

func (s *GetExtensionRequest) SetExtensionCode(v string) *GetExtensionRequest {
	s.ExtensionCode = &v
	return s
}

func (s *GetExtensionRequest) Validate() error {
	return dara.Validate(s)
}
