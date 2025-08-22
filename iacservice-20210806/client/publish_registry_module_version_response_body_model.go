// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRegistryModuleVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PublishRegistryModuleVersionResponseBody
	GetRequestId() *string
	SetVersion(v string) *PublishRegistryModuleVersionResponseBody
	GetVersion() *string
}

type PublishRegistryModuleVersionResponseBody struct {
	// example:
	//
	// 36E1679B-4D91-5AF6-B505-B5D4ACDF75BD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1.1.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s PublishRegistryModuleVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishRegistryModuleVersionResponseBody) GoString() string {
	return s.String()
}

func (s *PublishRegistryModuleVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishRegistryModuleVersionResponseBody) GetVersion() *string {
	return s.Version
}

func (s *PublishRegistryModuleVersionResponseBody) SetRequestId(v string) *PublishRegistryModuleVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishRegistryModuleVersionResponseBody) SetVersion(v string) *PublishRegistryModuleVersionResponseBody {
	s.Version = &v
	return s
}

func (s *PublishRegistryModuleVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
