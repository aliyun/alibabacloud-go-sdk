// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZnodeChildrenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListZnodeChildrenRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *ListZnodeChildrenRequest
	GetClusterId() *string
	SetPath(v string) *ListZnodeChildrenRequest
	GetPath() *string
}

type ListZnodeChildrenRequest struct {
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
}

func (s ListZnodeChildrenRequest) String() string {
	return dara.Prettify(s)
}

func (s ListZnodeChildrenRequest) GoString() string {
	return s.String()
}

func (s *ListZnodeChildrenRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListZnodeChildrenRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListZnodeChildrenRequest) GetPath() *string {
	return s.Path
}

func (s *ListZnodeChildrenRequest) SetAcceptLanguage(v string) *ListZnodeChildrenRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListZnodeChildrenRequest) SetClusterId(v string) *ListZnodeChildrenRequest {
	s.ClusterId = &v
	return s
}

func (s *ListZnodeChildrenRequest) SetPath(v string) *ListZnodeChildrenRequest {
	s.Path = &v
	return s
}

func (s *ListZnodeChildrenRequest) Validate() error {
	return dara.Validate(s)
}
