// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteZnodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteZnodeRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *DeleteZnodeRequest
	GetClusterId() *string
	SetPath(v string) *DeleteZnodeRequest
	GetPath() *string
	SetRequestPars(v string) *DeleteZnodeRequest
	GetRequestPars() *string
}

type DeleteZnodeRequest struct {
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

func (s DeleteZnodeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteZnodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteZnodeRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteZnodeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteZnodeRequest) GetPath() *string {
	return s.Path
}

func (s *DeleteZnodeRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *DeleteZnodeRequest) SetAcceptLanguage(v string) *DeleteZnodeRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteZnodeRequest) SetClusterId(v string) *DeleteZnodeRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteZnodeRequest) SetPath(v string) *DeleteZnodeRequest {
	s.Path = &v
	return s
}

func (s *DeleteZnodeRequest) SetRequestPars(v string) *DeleteZnodeRequest {
	s.RequestPars = &v
	return s
}

func (s *DeleteZnodeRequest) Validate() error {
	return dara.Validate(s)
}
