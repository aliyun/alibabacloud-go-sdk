// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZnodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateZnodeRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *UpdateZnodeRequest
	GetClusterId() *string
	SetData(v string) *UpdateZnodeRequest
	GetData() *string
	SetPath(v string) *UpdateZnodeRequest
	GetPath() *string
	SetRequestPars(v string) *UpdateZnodeRequest
	GetRequestPars() *string
}

type UpdateZnodeRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The data of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// data
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The path of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// /zookeeper
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s UpdateZnodeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateZnodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateZnodeRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateZnodeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateZnodeRequest) GetData() *string {
	return s.Data
}

func (s *UpdateZnodeRequest) GetPath() *string {
	return s.Path
}

func (s *UpdateZnodeRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *UpdateZnodeRequest) SetAcceptLanguage(v string) *UpdateZnodeRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateZnodeRequest) SetClusterId(v string) *UpdateZnodeRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateZnodeRequest) SetData(v string) *UpdateZnodeRequest {
	s.Data = &v
	return s
}

func (s *UpdateZnodeRequest) SetPath(v string) *UpdateZnodeRequest {
	s.Path = &v
	return s
}

func (s *UpdateZnodeRequest) SetRequestPars(v string) *UpdateZnodeRequest {
	s.RequestPars = &v
	return s
}

func (s *UpdateZnodeRequest) Validate() error {
	return dara.Validate(s)
}
