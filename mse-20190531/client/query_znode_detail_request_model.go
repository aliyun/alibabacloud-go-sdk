// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryZnodeDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryZnodeDetailRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *QueryZnodeDetailRequest
	GetClusterId() *string
	SetInstanceId(v string) *QueryZnodeDetailRequest
	GetInstanceId() *string
	SetPath(v string) *QueryZnodeDetailRequest
	GetPath() *string
	SetRequestPars(v string) *QueryZnodeDetailRequest
	GetRequestPars() *string
}

type QueryZnodeDetailRequest struct {
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
	// The ID of the cluster.
	//
	// example:
	//
	// mse-09k1q11****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-5bffa4e8630
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s QueryZnodeDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryZnodeDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryZnodeDetailRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryZnodeDetailRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryZnodeDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryZnodeDetailRequest) GetPath() *string {
	return s.Path
}

func (s *QueryZnodeDetailRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *QueryZnodeDetailRequest) SetAcceptLanguage(v string) *QueryZnodeDetailRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryZnodeDetailRequest) SetClusterId(v string) *QueryZnodeDetailRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryZnodeDetailRequest) SetInstanceId(v string) *QueryZnodeDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryZnodeDetailRequest) SetPath(v string) *QueryZnodeDetailRequest {
	s.Path = &v
	return s
}

func (s *QueryZnodeDetailRequest) SetRequestPars(v string) *QueryZnodeDetailRequest {
	s.RequestPars = &v
	return s
}

func (s *QueryZnodeDetailRequest) Validate() error {
	return dara.Validate(s)
}
