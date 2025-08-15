// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAvailableZonesResponseBody
	GetCode() *string
	SetData(v []*ListAvailableZonesResponseBodyData) *ListAvailableZonesResponseBody
	GetData() []*ListAvailableZonesResponseBodyData
	SetDynamicCode(v string) *ListAvailableZonesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAvailableZonesResponseBody
	GetDynamicMessage() *string
	SetHttpStatusCode(v int32) *ListAvailableZonesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAvailableZonesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAvailableZonesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAvailableZonesResponseBody
	GetSuccess() *bool
}

type ListAvailableZonesResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// Topic.NotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The result data that is returned.
	Data []*ListAvailableZonesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The dynamic error code.
	//
	// example:
	//
	// InstanceId
	DynamicCode *string `json:"dynamicCode,omitempty" xml:"dynamicCode,omitempty"`
	// The dynamic error message.
	//
	// example:
	//
	// InstanceId
	DynamicMessage *string `json:"dynamicMessage,omitempty" xml:"dynamicMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Parameter instanceId is mandatory for this action .
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the request. Each request has a unique ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AF9A8B10-C426-530F-A0DD-96320B39****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListAvailableZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableZonesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAvailableZonesResponseBody) GetData() []*ListAvailableZonesResponseBodyData {
	return s.Data
}

func (s *ListAvailableZonesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAvailableZonesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAvailableZonesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAvailableZonesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAvailableZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAvailableZonesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAvailableZonesResponseBody) SetCode(v string) *ListAvailableZonesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetData(v []*ListAvailableZonesResponseBodyData) *ListAvailableZonesResponseBody {
	s.Data = v
	return s
}

func (s *ListAvailableZonesResponseBody) SetDynamicCode(v string) *ListAvailableZonesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetDynamicMessage(v string) *ListAvailableZonesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetHttpStatusCode(v int32) *ListAvailableZonesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetMessage(v string) *ListAvailableZonesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetRequestId(v string) *ListAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableZonesResponseBody) SetSuccess(v bool) *ListAvailableZonesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAvailableZonesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAvailableZonesResponseBodyData struct {
	// The time when the zone was created.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the zone was last updated.
	//
	// example:
	//
	// 2022-08-01 20:05:50
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The ID of the current zone.
	//
	// example:
	//
	// cn-qingdao-b
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
	// The name of the current zone.
	//
	// example:
	//
	// ha-cn-t9b30w902vm_qrs
	ZoneName *string `json:"zoneName,omitempty" xml:"zoneName,omitempty"`
}

func (s ListAvailableZonesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableZonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAvailableZonesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAvailableZonesResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAvailableZonesResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListAvailableZonesResponseBodyData) GetZoneName() *string {
	return s.ZoneName
}

func (s *ListAvailableZonesResponseBodyData) SetCreateTime(v string) *ListAvailableZonesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAvailableZonesResponseBodyData) SetUpdateTime(v string) *ListAvailableZonesResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListAvailableZonesResponseBodyData) SetZoneId(v string) *ListAvailableZonesResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *ListAvailableZonesResponseBodyData) SetZoneName(v string) *ListAvailableZonesResponseBodyData {
	s.ZoneName = &v
	return s
}

func (s *ListAvailableZonesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
