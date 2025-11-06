// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayRouteOnAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGatewayRouteOnAuthResponseBody
	GetCode() *int32
	SetData(v []*ListGatewayRouteOnAuthResponseBodyData) *ListGatewayRouteOnAuthResponseBody
	GetData() []*ListGatewayRouteOnAuthResponseBodyData
	SetHttpStatusCode(v int32) *ListGatewayRouteOnAuthResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGatewayRouteOnAuthResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayRouteOnAuthResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGatewayRouteOnAuthResponseBody
	GetSuccess() *bool
}

type ListGatewayRouteOnAuthResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	Data []*ListGatewayRouteOnAuthResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 58E06A0A-BD2C-47A0-99C2-B100F353****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayRouteOnAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteOnAuthResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteOnAuthResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewayRouteOnAuthResponseBody) GetData() []*ListGatewayRouteOnAuthResponseBodyData {
	return s.Data
}

func (s *ListGatewayRouteOnAuthResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayRouteOnAuthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayRouteOnAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayRouteOnAuthResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayRouteOnAuthResponseBody) SetCode(v int32) *ListGatewayRouteOnAuthResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBody) SetData(v []*ListGatewayRouteOnAuthResponseBodyData) *ListGatewayRouteOnAuthResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBody) SetHttpStatusCode(v int32) *ListGatewayRouteOnAuthResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBody) SetMessage(v string) *ListGatewayRouteOnAuthResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBody) SetRequestId(v string) *ListGatewayRouteOnAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBody) SetSuccess(v bool) *ListGatewayRouteOnAuthResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBody) Validate() error {
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

type ListGatewayRouteOnAuthResponseBodyData struct {
	// The domain ID.
	//
	// example:
	//
	// 235
	DomainId *int64 `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The domain IDs.
	DomainIdList []*int64 `json:"DomainIdList,omitempty" xml:"DomainIdList,omitempty" type:"Repeated"`
	// The domain name.
	//
	// example:
	//
	// 123.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The domain names.
	DomainNameList []*string `json:"DomainNameList,omitempty" xml:"DomainNameList,omitempty" type:"Repeated"`
	// The gateway ID.
	//
	// example:
	//
	// 399
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597c****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The route ID.
	//
	// example:
	//
	// 12
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the route.
	//
	// example:
	//
	// ceshi
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The information about route matching.
	RoutePredicates *ListGatewayRouteOnAuthResponseBodyDataRoutePredicates `json:"RoutePredicates,omitempty" xml:"RoutePredicates,omitempty" type:"Struct"`
}

func (s ListGatewayRouteOnAuthResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteOnAuthResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteOnAuthResponseBodyData) GetDomainId() *int64 {
	return s.DomainId
}

func (s *ListGatewayRouteOnAuthResponseBodyData) GetDomainIdList() []*int64 {
	return s.DomainIdList
}

func (s *ListGatewayRouteOnAuthResponseBodyData) GetDomainName() *string {
	return s.DomainName
}

func (s *ListGatewayRouteOnAuthResponseBodyData) GetDomainNameList() []*string {
	return s.DomainNameList
}

func (s *ListGatewayRouteOnAuthResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewayRouteOnAuthResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListGatewayRouteOnAuthResponseBodyData) GetId() *int32 {
	return s.Id
}

func (s *ListGatewayRouteOnAuthResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListGatewayRouteOnAuthResponseBodyData) GetRoutePredicates() *ListGatewayRouteOnAuthResponseBodyDataRoutePredicates {
	return s.RoutePredicates
}

func (s *ListGatewayRouteOnAuthResponseBodyData) SetDomainId(v int64) *ListGatewayRouteOnAuthResponseBodyData {
	s.DomainId = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyData) SetDomainIdList(v []*int64) *ListGatewayRouteOnAuthResponseBodyData {
	s.DomainIdList = v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyData) SetDomainName(v string) *ListGatewayRouteOnAuthResponseBodyData {
	s.DomainName = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyData) SetDomainNameList(v []*string) *ListGatewayRouteOnAuthResponseBodyData {
	s.DomainNameList = v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyData) SetGatewayId(v string) *ListGatewayRouteOnAuthResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyData) SetGatewayUniqueId(v string) *ListGatewayRouteOnAuthResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyData) SetId(v int32) *ListGatewayRouteOnAuthResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyData) SetName(v string) *ListGatewayRouteOnAuthResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyData) SetRoutePredicates(v *ListGatewayRouteOnAuthResponseBodyDataRoutePredicates) *ListGatewayRouteOnAuthResponseBodyData {
	s.RoutePredicates = v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyData) Validate() error {
	if s.RoutePredicates != nil {
		if err := s.RoutePredicates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayRouteOnAuthResponseBodyDataRoutePredicates struct {
	// The information about route matching.
	PathPredicates *ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates `json:"PathPredicates,omitempty" xml:"PathPredicates,omitempty" type:"Struct"`
}

func (s ListGatewayRouteOnAuthResponseBodyDataRoutePredicates) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteOnAuthResponseBodyDataRoutePredicates) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteOnAuthResponseBodyDataRoutePredicates) GetPathPredicates() *ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates {
	return s.PathPredicates
}

func (s *ListGatewayRouteOnAuthResponseBodyDataRoutePredicates) SetPathPredicates(v *ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates) *ListGatewayRouteOnAuthResponseBodyDataRoutePredicates {
	s.PathPredicates = v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyDataRoutePredicates) Validate() error {
	if s.PathPredicates != nil {
		if err := s.PathPredicates.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates struct {
	// The path.
	//
	// example:
	//
	// /api
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The matching type.
	//
	// example:
	//
	// PRE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates) GoString() string {
	return s.String()
}

func (s *ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates) GetPath() *string {
	return s.Path
}

func (s *ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates) GetType() *string {
	return s.Type
}

func (s *ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates) SetPath(v string) *ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates {
	s.Path = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates) SetType(v string) *ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates {
	s.Type = &v
	return s
}

func (s *ListGatewayRouteOnAuthResponseBodyDataRoutePredicatesPathPredicates) Validate() error {
	return dara.Validate(s)
}
