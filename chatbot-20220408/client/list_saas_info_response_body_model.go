// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSaasInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListSaasInfoResponseBodyData) *ListSaasInfoResponseBody
	GetData() []*ListSaasInfoResponseBodyData
	SetRequestId(v string) *ListSaasInfoResponseBody
	GetRequestId() *string
	SetSaasToken(v string) *ListSaasInfoResponseBody
	GetSaasToken() *string
}

type ListSaasInfoResponseBody struct {
	Data []*ListSaasInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// A629A28F-F25E-5572-A679-FA46FB0151D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 06614fdb-c72f-436e-8003-dfe8a2854a15
	SaasToken *string `json:"SaasToken,omitempty" xml:"SaasToken,omitempty"`
}

func (s ListSaasInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSaasInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListSaasInfoResponseBody) GetData() []*ListSaasInfoResponseBodyData {
	return s.Data
}

func (s *ListSaasInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSaasInfoResponseBody) GetSaasToken() *string {
	return s.SaasToken
}

func (s *ListSaasInfoResponseBody) SetData(v []*ListSaasInfoResponseBodyData) *ListSaasInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListSaasInfoResponseBody) SetRequestId(v string) *ListSaasInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSaasInfoResponseBody) SetSaasToken(v string) *ListSaasInfoResponseBody {
	s.SaasToken = &v
	return s
}

func (s *ListSaasInfoResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSaasInfoResponseBodyData struct {
	// example:
	//
	// GLOBAL_SERVICE
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// GLOBAL SERVICE
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// https://pre-alime4service.console.aliyun.com/?productCode=p_beebot_public&switchAgent=1204001&saasCode=Robot&saasToken=06614fdb-c72f-436e-8003-dfe8a2854a15&saasName=123#/robot
	ServiceUrl *string `json:"ServiceUrl,omitempty" xml:"ServiceUrl,omitempty"`
	// example:
	//
	// https://alime.console.aliyun.com/?productCode=p_beebot_public&switchAgent=1204001&saasCode=Robot&saasToken=06614fdb-c72f-436e-8003-dfe8a2854a15&saasName=123#/robot
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListSaasInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSaasInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSaasInfoResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *ListSaasInfoResponseBodyData) GetEnName() *string {
	return s.EnName
}

func (s *ListSaasInfoResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListSaasInfoResponseBodyData) GetServiceUrl() *string {
	return s.ServiceUrl
}

func (s *ListSaasInfoResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *ListSaasInfoResponseBodyData) SetCode(v string) *ListSaasInfoResponseBodyData {
	s.Code = &v
	return s
}

func (s *ListSaasInfoResponseBodyData) SetEnName(v string) *ListSaasInfoResponseBodyData {
	s.EnName = &v
	return s
}

func (s *ListSaasInfoResponseBodyData) SetName(v string) *ListSaasInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSaasInfoResponseBodyData) SetServiceUrl(v string) *ListSaasInfoResponseBodyData {
	s.ServiceUrl = &v
	return s
}

func (s *ListSaasInfoResponseBodyData) SetUrl(v string) *ListSaasInfoResponseBodyData {
	s.Url = &v
	return s
}

func (s *ListSaasInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
