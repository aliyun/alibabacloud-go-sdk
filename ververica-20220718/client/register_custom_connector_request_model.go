// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterCustomConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJarUrl(v string) *RegisterCustomConnectorRequest
	GetJarUrl() *string
}

type RegisterCustomConnectorRequest struct {
	// The URL in which the JAR package of the custom connector is stored. The platform must be able to access this address.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://flink/connector/mysql123
	JarUrl *string `json:"jarUrl,omitempty" xml:"jarUrl,omitempty"`
}

func (s RegisterCustomConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterCustomConnectorRequest) GoString() string {
	return s.String()
}

func (s *RegisterCustomConnectorRequest) GetJarUrl() *string {
	return s.JarUrl
}

func (s *RegisterCustomConnectorRequest) SetJarUrl(v string) *RegisterCustomConnectorRequest {
	s.JarUrl = &v
	return s
}

func (s *RegisterCustomConnectorRequest) Validate() error {
	return dara.Validate(s)
}
