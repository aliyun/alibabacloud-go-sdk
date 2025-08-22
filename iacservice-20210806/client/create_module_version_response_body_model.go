// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModuleVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModuleVersion(v string) *CreateModuleVersionResponseBody
	GetModuleVersion() *string
	SetRequestId(v string) *CreateModuleVersionResponseBody
	GetRequestId() *string
}

type CreateModuleVersionResponseBody struct {
	// example:
	//
	// v1
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// B6ED9F71-7FA8-598E-B64D-4606FB3FCCC9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateModuleVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModuleVersionResponseBody) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *CreateModuleVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModuleVersionResponseBody) SetModuleVersion(v string) *CreateModuleVersionResponseBody {
	s.ModuleVersion = &v
	return s
}

func (s *CreateModuleVersionResponseBody) SetRequestId(v string) *CreateModuleVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModuleVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
