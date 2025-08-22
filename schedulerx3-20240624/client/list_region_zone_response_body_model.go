// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionZoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListRegionZoneResponseBody
	GetCode() *int32
	SetData(v []*ListRegionZoneResponseBodyData) *ListRegionZoneResponseBody
	GetData() []*ListRegionZoneResponseBodyData
	SetErrorCode(v string) *ListRegionZoneResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListRegionZoneResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRegionZoneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRegionZoneResponseBody
	GetSuccess() *bool
}

type ListRegionZoneResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*ListRegionZoneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// IllegalRequest
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 438737AC-760A-57D9-B646-B7EF79426243
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRegionZoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegionZoneResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionZoneResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListRegionZoneResponseBody) GetData() []*ListRegionZoneResponseBodyData {
	return s.Data
}

func (s *ListRegionZoneResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListRegionZoneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRegionZoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegionZoneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRegionZoneResponseBody) SetCode(v int32) *ListRegionZoneResponseBody {
	s.Code = &v
	return s
}

func (s *ListRegionZoneResponseBody) SetData(v []*ListRegionZoneResponseBodyData) *ListRegionZoneResponseBody {
	s.Data = v
	return s
}

func (s *ListRegionZoneResponseBody) SetErrorCode(v string) *ListRegionZoneResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListRegionZoneResponseBody) SetMessage(v string) *ListRegionZoneResponseBody {
	s.Message = &v
	return s
}

func (s *ListRegionZoneResponseBody) SetRequestId(v string) *ListRegionZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionZoneResponseBody) SetSuccess(v bool) *ListRegionZoneResponseBody {
	s.Success = &v
	return s
}

func (s *ListRegionZoneResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRegionZoneResponseBodyData struct {
	// example:
	//
	// E
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// zone id
	//
	// example:
	//
	// cn-beijing-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListRegionZoneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRegionZoneResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRegionZoneResponseBodyData) GetLocalName() *string {
	return s.LocalName
}

func (s *ListRegionZoneResponseBodyData) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListRegionZoneResponseBodyData) SetLocalName(v string) *ListRegionZoneResponseBodyData {
	s.LocalName = &v
	return s
}

func (s *ListRegionZoneResponseBodyData) SetZoneId(v string) *ListRegionZoneResponseBodyData {
	s.ZoneId = &v
	return s
}

func (s *ListRegionZoneResponseBodyData) Validate() error {
	return dara.Validate(s)
}
