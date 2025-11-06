// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGatewayRouteResponseBody
	GetCode() *int32
	SetData(v *DeleteGatewayRouteResponseBodyData) *DeleteGatewayRouteResponseBody
	GetData() *DeleteGatewayRouteResponseBodyData
	SetHttpStatusCode(v int32) *DeleteGatewayRouteResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteGatewayRouteResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGatewayRouteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGatewayRouteResponseBody
	GetSuccess() *bool
}

type DeleteGatewayRouteResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DeleteGatewayRouteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ADDD8AB7-8D1C-4697-A83E-410D2607****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGatewayRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRouteResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGatewayRouteResponseBody) GetData() *DeleteGatewayRouteResponseBodyData {
	return s.Data
}

func (s *DeleteGatewayRouteResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteGatewayRouteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGatewayRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGatewayRouteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGatewayRouteResponseBody) SetCode(v int32) *DeleteGatewayRouteResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGatewayRouteResponseBody) SetData(v *DeleteGatewayRouteResponseBodyData) *DeleteGatewayRouteResponseBody {
	s.Data = v
	return s
}

func (s *DeleteGatewayRouteResponseBody) SetHttpStatusCode(v int32) *DeleteGatewayRouteResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGatewayRouteResponseBody) SetMessage(v string) *DeleteGatewayRouteResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGatewayRouteResponseBody) SetRequestId(v string) *DeleteGatewayRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGatewayRouteResponseBody) SetSuccess(v bool) *DeleteGatewayRouteResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGatewayRouteResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteGatewayRouteResponseBodyData struct {
	// The default service ID.
	//
	// example:
	//
	// 1
	DefaultServiceId *int64 `json:"DefaultServiceId,omitempty" xml:"DefaultServiceId,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 125
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-5017305290e14cebbrvec4a5****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-07 18:07:57
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2021-12-30T06:41:52.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID.
	//
	// example:
	//
	// 12
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The matching rules.
	//
	// example:
	//
	// {\\"PathPredicates\\":{\\"Path\\":\\"/metas\\",\\"Type\\":\\"PRE\\",\\"IgnoreCase\\":false}}
	Predicates *string `json:"Predicates,omitempty" xml:"Predicates,omitempty"`
	// The sequence number of the route.
	//
	// example:
	//
	// 1
	RouteOrder *int32 `json:"RouteOrder,omitempty" xml:"RouteOrder,omitempty"`
	// The status of the route. Valid values:
	//
	// 	- 0: unpublished
	//
	// 	- 2: publishing
	//
	// 	- 3: published
	//
	// 	- 4: editing (updated but not published)
	//
	// 	- 5: unpublishing
	//
	// 	- 6: unavailable
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteGatewayRouteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayRouteResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteGatewayRouteResponseBodyData) GetDefaultServiceId() *int64 {
	return s.DefaultServiceId
}

func (s *DeleteGatewayRouteResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *DeleteGatewayRouteResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayRouteResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DeleteGatewayRouteResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DeleteGatewayRouteResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *DeleteGatewayRouteResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DeleteGatewayRouteResponseBodyData) GetPredicates() *string {
	return s.Predicates
}

func (s *DeleteGatewayRouteResponseBodyData) GetRouteOrder() *int32 {
	return s.RouteOrder
}

func (s *DeleteGatewayRouteResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *DeleteGatewayRouteResponseBodyData) SetDefaultServiceId(v int64) *DeleteGatewayRouteResponseBodyData {
	s.DefaultServiceId = &v
	return s
}

func (s *DeleteGatewayRouteResponseBodyData) SetGatewayId(v int64) *DeleteGatewayRouteResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *DeleteGatewayRouteResponseBodyData) SetGatewayUniqueId(v string) *DeleteGatewayRouteResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayRouteResponseBodyData) SetGmtCreate(v string) *DeleteGatewayRouteResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *DeleteGatewayRouteResponseBodyData) SetGmtModified(v string) *DeleteGatewayRouteResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DeleteGatewayRouteResponseBodyData) SetId(v int64) *DeleteGatewayRouteResponseBodyData {
	s.Id = &v
	return s
}

func (s *DeleteGatewayRouteResponseBodyData) SetName(v string) *DeleteGatewayRouteResponseBodyData {
	s.Name = &v
	return s
}

func (s *DeleteGatewayRouteResponseBodyData) SetPredicates(v string) *DeleteGatewayRouteResponseBodyData {
	s.Predicates = &v
	return s
}

func (s *DeleteGatewayRouteResponseBodyData) SetRouteOrder(v int32) *DeleteGatewayRouteResponseBodyData {
	s.RouteOrder = &v
	return s
}

func (s *DeleteGatewayRouteResponseBodyData) SetStatus(v int32) *DeleteGatewayRouteResponseBodyData {
	s.Status = &v
	return s
}

func (s *DeleteGatewayRouteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
