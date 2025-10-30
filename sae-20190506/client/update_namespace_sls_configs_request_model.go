// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceSlsConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNameSpaceShortId(v string) *UpdateNamespaceSlsConfigsRequest
	GetNameSpaceShortId() *string
	SetNamespaceId(v string) *UpdateNamespaceSlsConfigsRequest
	GetNamespaceId() *string
	SetSlsConfigs(v string) *UpdateNamespaceSlsConfigsRequest
	GetSlsConfigs() *string
}

type UpdateNamespaceSlsConfigsRequest struct {
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// [{"projectName":"test","logDir":"/tmp/a.log","logstoreName":"sae","logtailName":""}]
	SlsConfigs *string `json:"SlsConfigs,omitempty" xml:"SlsConfigs,omitempty"`
}

func (s UpdateNamespaceSlsConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceSlsConfigsRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceSlsConfigsRequest) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *UpdateNamespaceSlsConfigsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateNamespaceSlsConfigsRequest) GetSlsConfigs() *string {
	return s.SlsConfigs
}

func (s *UpdateNamespaceSlsConfigsRequest) SetNameSpaceShortId(v string) *UpdateNamespaceSlsConfigsRequest {
	s.NameSpaceShortId = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsRequest) SetNamespaceId(v string) *UpdateNamespaceSlsConfigsRequest {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsRequest) SetSlsConfigs(v string) *UpdateNamespaceSlsConfigsRequest {
	s.SlsConfigs = &v
	return s
}

func (s *UpdateNamespaceSlsConfigsRequest) Validate() error {
	return dara.Validate(s)
}
