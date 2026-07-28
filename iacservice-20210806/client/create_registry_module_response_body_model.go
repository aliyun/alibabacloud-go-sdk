// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRegistryModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRegistryModuleResponseBody
	GetRequestId() *string
	SetSource(v string) *CreateRegistryModuleResponseBody
	GetSource() *string
}

type CreateRegistryModuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C896FE0A-1BEA-5D01-BFF4-B03B82B9CA3D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The template source, which is a concatenation of \\<namespaceName>/\\<ModuleName>.
	//
	// example:
	//
	// namespaceName/ModuleName
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s CreateRegistryModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRegistryModuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRegistryModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRegistryModuleResponseBody) GetSource() *string {
	return s.Source
}

func (s *CreateRegistryModuleResponseBody) SetRequestId(v string) *CreateRegistryModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRegistryModuleResponseBody) SetSource(v string) *CreateRegistryModuleResponseBody {
	s.Source = &v
	return s
}

func (s *CreateRegistryModuleResponseBody) Validate() error {
	return dara.Validate(s)
}
