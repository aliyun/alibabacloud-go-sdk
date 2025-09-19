// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleConfigStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModuleNames(v []*string) *GetModuleConfigStatusRequest
	GetModuleNames() []*string
}

type GetModuleConfigStatusRequest struct {
	// The service modules that you want to query.
	//
	// This parameter is required.
	ModuleNames []*string `json:"ModuleNames,omitempty" xml:"ModuleNames,omitempty" type:"Repeated"`
}

func (s GetModuleConfigStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModuleConfigStatusRequest) GoString() string {
	return s.String()
}

func (s *GetModuleConfigStatusRequest) GetModuleNames() []*string {
	return s.ModuleNames
}

func (s *GetModuleConfigStatusRequest) SetModuleNames(v []*string) *GetModuleConfigStatusRequest {
	s.ModuleNames = v
	return s
}

func (s *GetModuleConfigStatusRequest) Validate() error {
	return dara.Validate(s)
}
