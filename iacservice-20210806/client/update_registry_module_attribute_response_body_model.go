// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRegistryModuleAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRegistryModuleAttributeResponseBody
	GetRequestId() *string
}

type UpdateRegistryModuleAttributeResponseBody struct {
	// example:
	//
	// 127A5B81-D1E7-5E33-8D44-B89507C4B81F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateRegistryModuleAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRegistryModuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRegistryModuleAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRegistryModuleAttributeResponseBody) SetRequestId(v string) *UpdateRegistryModuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRegistryModuleAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
