// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayWebsocketUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGatewayWebsocketUsageResponseBody
	GetCode() *string
	SetData(v *GetGatewayWebsocketUsageResponseBodyData) *GetGatewayWebsocketUsageResponseBody
	GetData() *GetGatewayWebsocketUsageResponseBodyData
	SetMessage(v string) *GetGatewayWebsocketUsageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayWebsocketUsageResponseBody
	GetRequestId() *string
}

type GetGatewayWebsocketUsageResponseBody struct {
	// The business response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The query result.
	//
	// example:
	//
	// {}
	Data *GetGatewayWebsocketUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// The specified gateway does not exist.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D0A6A1A0-8793-4C10-AB1A-2C03E770A912
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayWebsocketUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayWebsocketUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayWebsocketUsageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGatewayWebsocketUsageResponseBody) GetData() *GetGatewayWebsocketUsageResponseBodyData {
	return s.Data
}

func (s *GetGatewayWebsocketUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayWebsocketUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayWebsocketUsageResponseBody) SetCode(v string) *GetGatewayWebsocketUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayWebsocketUsageResponseBody) SetData(v *GetGatewayWebsocketUsageResponseBodyData) *GetGatewayWebsocketUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayWebsocketUsageResponseBody) SetMessage(v string) *GetGatewayWebsocketUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayWebsocketUsageResponseBody) SetRequestId(v string) *GetGatewayWebsocketUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayWebsocketUsageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayWebsocketUsageResponseBodyData struct {
	// The data basis.
	//
	// example:
	//
	// redis_period
	Basis *string `json:"basis,omitempty" xml:"basis,omitempty"`
	// Indicates whether a next page is available.
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The cycle usage.
	//
	// example:
	//
	// []
	Items []map[string]interface{} `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s GetGatewayWebsocketUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayWebsocketUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayWebsocketUsageResponseBodyData) GetBasis() *string {
	return s.Basis
}

func (s *GetGatewayWebsocketUsageResponseBodyData) GetHasMore() *bool {
	return s.HasMore
}

func (s *GetGatewayWebsocketUsageResponseBodyData) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *GetGatewayWebsocketUsageResponseBodyData) SetBasis(v string) *GetGatewayWebsocketUsageResponseBodyData {
	s.Basis = &v
	return s
}

func (s *GetGatewayWebsocketUsageResponseBodyData) SetHasMore(v bool) *GetGatewayWebsocketUsageResponseBodyData {
	s.HasMore = &v
	return s
}

func (s *GetGatewayWebsocketUsageResponseBodyData) SetItems(v []map[string]interface{}) *GetGatewayWebsocketUsageResponseBodyData {
	s.Items = v
	return s
}

func (s *GetGatewayWebsocketUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
