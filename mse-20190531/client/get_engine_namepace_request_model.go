// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEngineNamepaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *GetEngineNamepaceRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *GetEngineNamepaceRequest
	GetClusterId() *string
	SetId(v string) *GetEngineNamepaceRequest
	GetId() *string
	SetInstanceId(v string) *GetEngineNamepaceRequest
	GetInstanceId() *string
}

type GetEngineNamepaceRequest struct {
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
	// mse-98s****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The destination ID.
	//
	// example:
	//
	// 0e958d79-****-b282-b702d66362b5
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse.cn-hangzhou.aliyuncs.com
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetEngineNamepaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEngineNamepaceRequest) GoString() string {
	return s.String()
}

func (s *GetEngineNamepaceRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *GetEngineNamepaceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetEngineNamepaceRequest) GetId() *string {
	return s.Id
}

func (s *GetEngineNamepaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetEngineNamepaceRequest) SetAcceptLanguage(v string) *GetEngineNamepaceRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetEngineNamepaceRequest) SetClusterId(v string) *GetEngineNamepaceRequest {
	s.ClusterId = &v
	return s
}

func (s *GetEngineNamepaceRequest) SetId(v string) *GetEngineNamepaceRequest {
	s.Id = &v
	return s
}

func (s *GetEngineNamepaceRequest) SetInstanceId(v string) *GetEngineNamepaceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetEngineNamepaceRequest) Validate() error {
	return dara.Validate(s)
}
