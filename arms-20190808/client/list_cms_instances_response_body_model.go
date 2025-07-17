// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCmsInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCmsInstancesResponseBodyData) *ListCmsInstancesResponseBody
	GetData() *ListCmsInstancesResponseBodyData
	SetRequestId(v string) *ListCmsInstancesResponseBody
	GetRequestId() *string
}

type ListCmsInstancesResponseBody struct {
	Data *ListCmsInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// E7A04B0D-E2CA-59BB-8A9D-D5D349C22BF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCmsInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCmsInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCmsInstancesResponseBody) GetData() *ListCmsInstancesResponseBodyData {
	return s.Data
}

func (s *ListCmsInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCmsInstancesResponseBody) SetData(v *ListCmsInstancesResponseBodyData) *ListCmsInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListCmsInstancesResponseBody) SetRequestId(v string) *ListCmsInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCmsInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCmsInstancesResponseBodyData struct {
	// example:
	//
	// true
	EnableTag *bool                                       `json:"EnableTag,omitempty" xml:"EnableTag,omitempty"`
	Products  []*ListCmsInstancesResponseBodyDataProducts `json:"Products,omitempty" xml:"Products,omitempty" type:"Repeated"`
}

func (s ListCmsInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCmsInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCmsInstancesResponseBodyData) GetEnableTag() *bool {
	return s.EnableTag
}

func (s *ListCmsInstancesResponseBodyData) GetProducts() []*ListCmsInstancesResponseBodyDataProducts {
	return s.Products
}

func (s *ListCmsInstancesResponseBodyData) SetEnableTag(v bool) *ListCmsInstancesResponseBodyData {
	s.EnableTag = &v
	return s
}

func (s *ListCmsInstancesResponseBodyData) SetProducts(v []*ListCmsInstancesResponseBodyDataProducts) *ListCmsInstancesResponseBodyData {
	s.Products = v
	return s
}

func (s *ListCmsInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListCmsInstancesResponseBodyDataProducts struct {
	// example:
	//
	// -
	Descr *string `json:"Descr,omitempty" xml:"Descr,omitempty"`
	// example:
	//
	// 20210
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// cloudserver
	Instance *string `json:"Instance,omitempty" xml:"Instance,omitempty"`
	// example:
	//
	// hologres
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// face
	Prod *string `json:"Prod,omitempty" xml:"Prod,omitempty"`
	// example:
	//
	// arms
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// true
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// 1647852021000
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// HOLOGRES
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// https://g.console.aliyun.com/d/1098370038733503-35894-565/cms-hologres?orgId\\u003d9\\u0026refresh\\u003d60s
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListCmsInstancesResponseBodyDataProducts) String() string {
	return dara.Prettify(s)
}

func (s ListCmsInstancesResponseBodyDataProducts) GoString() string {
	return s.String()
}

func (s *ListCmsInstancesResponseBodyDataProducts) GetDescr() *string {
	return s.Descr
}

func (s *ListCmsInstancesResponseBodyDataProducts) GetId() *string {
	return s.Id
}

func (s *ListCmsInstancesResponseBodyDataProducts) GetInstance() *string {
	return s.Instance
}

func (s *ListCmsInstancesResponseBodyDataProducts) GetName() *string {
	return s.Name
}

func (s *ListCmsInstancesResponseBodyDataProducts) GetProd() *string {
	return s.Prod
}

func (s *ListCmsInstancesResponseBodyDataProducts) GetSource() *string {
	return s.Source
}

func (s *ListCmsInstancesResponseBodyDataProducts) GetState() *string {
	return s.State
}

func (s *ListCmsInstancesResponseBodyDataProducts) GetTime() *string {
	return s.Time
}

func (s *ListCmsInstancesResponseBodyDataProducts) GetType() *string {
	return s.Type
}

func (s *ListCmsInstancesResponseBodyDataProducts) GetUrl() *string {
	return s.Url
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetDescr(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Descr = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetId(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Id = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetInstance(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Instance = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetName(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Name = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetProd(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Prod = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetSource(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Source = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetState(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.State = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetTime(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Time = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetType(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Type = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) SetUrl(v string) *ListCmsInstancesResponseBodyDataProducts {
	s.Url = &v
	return s
}

func (s *ListCmsInstancesResponseBodyDataProducts) Validate() error {
	return dara.Validate(s)
}
