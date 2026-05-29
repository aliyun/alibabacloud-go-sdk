// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProviderEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListModelProviderEndpointsResponseBodyData) *ListModelProviderEndpointsResponseBody
	GetData() []*ListModelProviderEndpointsResponseBodyData
	SetRequestId(v string) *ListModelProviderEndpointsResponseBody
	GetRequestId() *string
}

type ListModelProviderEndpointsResponseBody struct {
	Data []*ListModelProviderEndpointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListModelProviderEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelProviderEndpointsResponseBody) GetData() []*ListModelProviderEndpointsResponseBodyData {
	return s.Data
}

func (s *ListModelProviderEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelProviderEndpointsResponseBody) SetData(v []*ListModelProviderEndpointsResponseBodyData) *ListModelProviderEndpointsResponseBody {
	s.Data = v
	return s
}

func (s *ListModelProviderEndpointsResponseBody) SetRequestId(v string) *ListModelProviderEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelProviderEndpointsResponseBody) Validate() error {
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

type ListModelProviderEndpointsResponseBodyData struct {
	Description  *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Endpoints    []*ListModelProviderEndpointsResponseBodyDataEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	ProviderName *string                                                `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	ProviderUrl  *string                                                `json:"ProviderUrl,omitempty" xml:"ProviderUrl,omitempty"`
}

func (s ListModelProviderEndpointsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderEndpointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModelProviderEndpointsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListModelProviderEndpointsResponseBodyData) GetEndpoints() []*ListModelProviderEndpointsResponseBodyDataEndpoints {
	return s.Endpoints
}

func (s *ListModelProviderEndpointsResponseBodyData) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListModelProviderEndpointsResponseBodyData) GetProviderUrl() *string {
	return s.ProviderUrl
}

func (s *ListModelProviderEndpointsResponseBodyData) SetDescription(v string) *ListModelProviderEndpointsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListModelProviderEndpointsResponseBodyData) SetEndpoints(v []*ListModelProviderEndpointsResponseBodyDataEndpoints) *ListModelProviderEndpointsResponseBodyData {
	s.Endpoints = v
	return s
}

func (s *ListModelProviderEndpointsResponseBodyData) SetProviderName(v string) *ListModelProviderEndpointsResponseBodyData {
	s.ProviderName = &v
	return s
}

func (s *ListModelProviderEndpointsResponseBodyData) SetProviderUrl(v string) *ListModelProviderEndpointsResponseBodyData {
	s.ProviderUrl = &v
	return s
}

func (s *ListModelProviderEndpointsResponseBodyData) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelProviderEndpointsResponseBodyDataEndpoints struct {
	ApiType     *string   `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	BaseUrl     *string   `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	ProviderUrl *string   `json:"ProviderUrl,omitempty" xml:"ProviderUrl,omitempty"`
	Tags        []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListModelProviderEndpointsResponseBodyDataEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderEndpointsResponseBodyDataEndpoints) GoString() string {
	return s.String()
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) GetApiType() *string {
	return s.ApiType
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) GetDescription() *string {
	return s.Description
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) GetName() *string {
	return s.Name
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) GetProviderUrl() *string {
	return s.ProviderUrl
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) GetTags() []*string {
	return s.Tags
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) SetApiType(v string) *ListModelProviderEndpointsResponseBodyDataEndpoints {
	s.ApiType = &v
	return s
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) SetBaseUrl(v string) *ListModelProviderEndpointsResponseBodyDataEndpoints {
	s.BaseUrl = &v
	return s
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) SetDescription(v string) *ListModelProviderEndpointsResponseBodyDataEndpoints {
	s.Description = &v
	return s
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) SetName(v string) *ListModelProviderEndpointsResponseBodyDataEndpoints {
	s.Name = &v
	return s
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) SetProviderUrl(v string) *ListModelProviderEndpointsResponseBodyDataEndpoints {
	s.ProviderUrl = &v
	return s
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) SetTags(v []*string) *ListModelProviderEndpointsResponseBodyDataEndpoints {
	s.Tags = v
	return s
}

func (s *ListModelProviderEndpointsResponseBodyDataEndpoints) Validate() error {
	return dara.Validate(s)
}
