// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgAndFactoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOrgAndFactoryResponseBody
	GetCode() *string
	SetData(v []*GetOrgAndFactoryResponseBodyData) *GetOrgAndFactoryResponseBody
	GetData() []*GetOrgAndFactoryResponseBodyData
	SetHttpCode(v int32) *GetOrgAndFactoryResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetOrgAndFactoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOrgAndFactoryResponseBody
	GetSuccess() *bool
}

type GetOrgAndFactoryResponseBody struct {
	// The code returned for the request.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// data
	Data []*GetOrgAndFactoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetOrgAndFactoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrgAndFactoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOrgAndFactoryResponseBody) GetData() []*GetOrgAndFactoryResponseBodyData {
	return s.Data
}

func (s *GetOrgAndFactoryResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetOrgAndFactoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrgAndFactoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOrgAndFactoryResponseBody) SetCode(v string) *GetOrgAndFactoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetData(v []*GetOrgAndFactoryResponseBodyData) *GetOrgAndFactoryResponseBody {
	s.Data = v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetHttpCode(v int32) *GetOrgAndFactoryResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetRequestId(v string) *GetOrgAndFactoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrgAndFactoryResponseBody) SetSuccess(v bool) *GetOrgAndFactoryResponseBody {
	s.Success = &v
	return s
}

func (s *GetOrgAndFactoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOrgAndFactoryResponseBodyData struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 1319617584664960
	AliyunPk *string `json:"aliyunPk,omitempty" xml:"aliyunPk,omitempty"`
	// The sites.
	FactoryList []*GetOrgAndFactoryResponseBodyDataFactoryList `json:"factoryList,omitempty" xml:"factoryList,omitempty" type:"Repeated"`
	// The enterprise ID.
	//
	// example:
	//
	// 6265f42XXXX2fec150
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// The enterprise name.
	//
	// example:
	//
	// Ledi Industrial Park
	OrganizationName *string `json:"organizationName,omitempty" xml:"organizationName,omitempty"`
}

func (s GetOrgAndFactoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOrgAndFactoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponseBodyData) GetAliyunPk() *string {
	return s.AliyunPk
}

func (s *GetOrgAndFactoryResponseBodyData) GetFactoryList() []*GetOrgAndFactoryResponseBodyDataFactoryList {
	return s.FactoryList
}

func (s *GetOrgAndFactoryResponseBodyData) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetOrgAndFactoryResponseBodyData) GetOrganizationName() *string {
	return s.OrganizationName
}

func (s *GetOrgAndFactoryResponseBodyData) SetAliyunPk(v string) *GetOrgAndFactoryResponseBodyData {
	s.AliyunPk = &v
	return s
}

func (s *GetOrgAndFactoryResponseBodyData) SetFactoryList(v []*GetOrgAndFactoryResponseBodyDataFactoryList) *GetOrgAndFactoryResponseBodyData {
	s.FactoryList = v
	return s
}

func (s *GetOrgAndFactoryResponseBodyData) SetOrganizationId(v string) *GetOrgAndFactoryResponseBodyData {
	s.OrganizationId = &v
	return s
}

func (s *GetOrgAndFactoryResponseBodyData) SetOrganizationName(v string) *GetOrgAndFactoryResponseBodyData {
	s.OrganizationName = &v
	return s
}

func (s *GetOrgAndFactoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetOrgAndFactoryResponseBodyDataFactoryList struct {
	// The site ID.
	//
	// example:
	//
	// pn_95
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	// The site name.
	//
	// example:
	//
	// Ledi Industrial Park 1
	FactoryName *string `json:"factoryName,omitempty" xml:"factoryName,omitempty"`
}

func (s GetOrgAndFactoryResponseBodyDataFactoryList) String() string {
	return dara.Prettify(s)
}

func (s GetOrgAndFactoryResponseBodyDataFactoryList) GoString() string {
	return s.String()
}

func (s *GetOrgAndFactoryResponseBodyDataFactoryList) GetFactoryId() *string {
	return s.FactoryId
}

func (s *GetOrgAndFactoryResponseBodyDataFactoryList) GetFactoryName() *string {
	return s.FactoryName
}

func (s *GetOrgAndFactoryResponseBodyDataFactoryList) SetFactoryId(v string) *GetOrgAndFactoryResponseBodyDataFactoryList {
	s.FactoryId = &v
	return s
}

func (s *GetOrgAndFactoryResponseBodyDataFactoryList) SetFactoryName(v string) *GetOrgAndFactoryResponseBodyDataFactoryList {
	s.FactoryName = &v
	return s
}

func (s *GetOrgAndFactoryResponseBodyDataFactoryList) Validate() error {
	return dara.Validate(s)
}
