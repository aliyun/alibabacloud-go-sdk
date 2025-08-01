// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateZnodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateZnodeRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *CreateZnodeRequest
	GetClusterId() *string
	SetData(v string) *CreateZnodeRequest
	GetData() *string
	SetPath(v string) *CreateZnodeRequest
	GetPath() *string
}

type CreateZnodeRequest struct {
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
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The data of the node.
	//
	// example:
	//
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The path of the node.
	//
	// example:
	//
	// /zookeeper
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s CreateZnodeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateZnodeRequest) GoString() string {
	return s.String()
}

func (s *CreateZnodeRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateZnodeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateZnodeRequest) GetData() *string {
	return s.Data
}

func (s *CreateZnodeRequest) GetPath() *string {
	return s.Path
}

func (s *CreateZnodeRequest) SetAcceptLanguage(v string) *CreateZnodeRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateZnodeRequest) SetClusterId(v string) *CreateZnodeRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateZnodeRequest) SetData(v string) *CreateZnodeRequest {
	s.Data = &v
	return s
}

func (s *CreateZnodeRequest) SetPath(v string) *CreateZnodeRequest {
	s.Path = &v
	return s
}

func (s *CreateZnodeRequest) Validate() error {
	return dara.Validate(s)
}
