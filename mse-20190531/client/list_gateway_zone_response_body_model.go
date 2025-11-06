// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListGatewayZoneResponseBody
	GetCode() *int32
	SetData(v []*ListGatewayZoneResponseBodyData) *ListGatewayZoneResponseBody
	GetData() []*ListGatewayZoneResponseBodyData
	SetDynamicCode(v string) *ListGatewayZoneResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListGatewayZoneResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *ListGatewayZoneResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ListGatewayZoneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGatewayZoneResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGatewayZoneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGatewayZoneResponseBody
	GetSuccess() *bool
}

type ListGatewayZoneResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried zones.
	Data []*ListGatewayZoneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic part in the error message. This parameter is used to replace %s in **ErrMessage**.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The status code.
	//
	// example:
	//
	// NO_PERMISSION
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EE5C32A1-BC0E-4B79-817C-103E4EDF****
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

func (s ListGatewayZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayZoneResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayZoneResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListGatewayZoneResponseBody) GetData() []*ListGatewayZoneResponseBodyData {
	return s.Data
}

func (s *ListGatewayZoneResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListGatewayZoneResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListGatewayZoneResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListGatewayZoneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayZoneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayZoneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayZoneResponseBody) SetCode(v int32) *ListGatewayZoneResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayZoneResponseBody) SetData(v []*ListGatewayZoneResponseBodyData) *ListGatewayZoneResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayZoneResponseBody) SetDynamicCode(v string) *ListGatewayZoneResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListGatewayZoneResponseBody) SetDynamicMessage(v string) *ListGatewayZoneResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListGatewayZoneResponseBody) SetErrorCode(v string) *ListGatewayZoneResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListGatewayZoneResponseBody) SetHttpStatusCode(v int32) *ListGatewayZoneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayZoneResponseBody) SetMessage(v string) *ListGatewayZoneResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayZoneResponseBody) SetRequestId(v string) *ListGatewayZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayZoneResponseBody) SetSuccess(v bool) *ListGatewayZoneResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayZoneResponseBody) Validate() error {
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

type ListGatewayZoneResponseBodyData struct {
	// The local name of the zone.
	//
	// example:
	//
	// I
	LocalName  *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	SupportQat *bool   `json:"SupportQat,omitempty" xml:"SupportQat,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListGatewayZoneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayZoneResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayZoneResponseBodyData) GetLocalName() *string {
	return s.LocalName
}

func (s *ListGatewayZoneResponseBodyData) GetSupportQat() *bool {
	return s.SupportQat
}

func (s *ListGatewayZoneResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListGatewayZoneResponseBodyData) SetLocalName(v string) *ListGatewayZoneResponseBodyData {
	s.LocalName = &v
	return s
}

func (s *ListGatewayZoneResponseBodyData) SetSupportQat(v bool) *ListGatewayZoneResponseBodyData {
	s.SupportQat = &v
	return s
}

func (s *ListGatewayZoneResponseBodyData) SetZoneId(v string) *ListGatewayZoneResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *ListGatewayZoneResponseBodyData) Validate() error {
	return dara.Validate(s)
}
