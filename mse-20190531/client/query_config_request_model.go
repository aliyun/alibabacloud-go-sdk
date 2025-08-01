// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryConfigRequest
	GetAcceptLanguage() *string
	SetClusterId(v string) *QueryConfigRequest
	GetClusterId() *string
	SetConfigType(v string) *QueryConfigRequest
	GetConfigType() *string
	SetInstanceId(v string) *QueryConfigRequest
	GetInstanceId() *string
	SetNeedRunningConf(v bool) *QueryConfigRequest
	GetNeedRunningConf() *bool
	SetRequestPars(v string) *QueryConfigRequest
	GetRequestPars() *string
}

type QueryConfigRequest struct {
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
	// A reserved parameter.
	//
	// example:
	//
	// null
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse_prepaid_public_cn-st2212****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether runtime configurations are required.
	//
	// example:
	//
	// true
	NeedRunningConf *bool `json:"NeedRunningConf,omitempty" xml:"NeedRunningConf,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s QueryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *QueryConfigRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *QueryConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryConfigRequest) GetNeedRunningConf() *bool {
	return s.NeedRunningConf
}

func (s *QueryConfigRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *QueryConfigRequest) SetAcceptLanguage(v string) *QueryConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryConfigRequest) SetClusterId(v string) *QueryConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *QueryConfigRequest) SetConfigType(v string) *QueryConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *QueryConfigRequest) SetInstanceId(v string) *QueryConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryConfigRequest) SetNeedRunningConf(v bool) *QueryConfigRequest {
	s.NeedRunningConf = &v
	return s
}

func (s *QueryConfigRequest) SetRequestPars(v string) *QueryConfigRequest {
	s.RequestPars = &v
	return s
}

func (s *QueryConfigRequest) Validate() error {
	return dara.Validate(s)
}
