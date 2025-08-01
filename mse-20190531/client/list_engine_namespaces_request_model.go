// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEngineNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListEngineNamespacesRequest
	GetAcceptLanguage() *string
	SetInstanceId(v string) *ListEngineNamespacesRequest
	GetInstanceId() *string
}

type ListEngineNamespacesRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-st21ri2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListEngineNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEngineNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListEngineNamespacesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListEngineNamespacesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEngineNamespacesRequest) SetAcceptLanguage(v string) *ListEngineNamespacesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListEngineNamespacesRequest) SetInstanceId(v string) *ListEngineNamespacesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListEngineNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
