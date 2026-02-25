// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListGatewayResponseBodyData) *ListGatewayResponseBody
	GetData() []*ListGatewayResponseBodyData
	SetErrCode(v string) *ListGatewayResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListGatewayResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ListGatewayResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListGatewayResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGatewayResponseBody
	GetSuccess() *bool
}

type ListGatewayResponseBody struct {
	Data []*ListGatewayResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// InvalidParams
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBody) GetData() []*ListGatewayResponseBodyData {
	return s.Data
}

func (s *ListGatewayResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListGatewayResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListGatewayResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGatewayResponseBody) SetData(v []*ListGatewayResponseBodyData) *ListGatewayResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayResponseBody) SetErrCode(v string) *ListGatewayResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListGatewayResponseBody) SetErrMessage(v string) *ListGatewayResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListGatewayResponseBody) SetHttpStatusCode(v int32) *ListGatewayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGatewayResponseBody) SetRequestId(v string) *ListGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayResponseBody) SetSuccess(v bool) *ListGatewayResponseBody {
	s.Success = &v
	return s
}

func (s *ListGatewayResponseBody) Validate() error {
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

type ListGatewayResponseBodyData struct {
	// example:
	//
	// true
	EnablePublicNet *bool `json:"EnablePublicNet,omitempty" xml:"EnablePublicNet,omitempty"`
	// example:
	//
	// 2
	FeNodeNumber *int32 `json:"FeNodeNumber,omitempty" xml:"FeNodeNumber,omitempty"`
	// example:
	//
	// 13822
	GatewayId   *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// example:
	//
	// slb
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// example:
	//
	// fe-c-b25e21e24388****-ab8sjd-internal.starrocks.aliyuncs.com
	InternalDomain *string `json:"InternalDomain,omitempty" xml:"InternalDomain,omitempty"`
	// example:
	//
	// lb-123abc
	InternalSlbId *string `json:"InternalSlbId,omitempty" xml:"InternalSlbId,omitempty"`
	// PrivatezoneId
	//
	// example:
	//
	// a62des2123243881b9s2sa220k2l38m9
	PrivatezoneId *string `json:"PrivatezoneId,omitempty" xml:"PrivatezoneId,omitempty"`
	// example:
	//
	// fe-c-b25e21e24388****-8s272d.starrocks.aliyuncs.com
	PublicDomain *string `json:"PublicDomain,omitempty" xml:"PublicDomain,omitempty"`
	// example:
	//
	// acl-hsb123ksi2
	PublicSlbAclId *string `json:"PublicSlbAclId,omitempty" xml:"PublicSlbAclId,omitempty"`
	// example:
	//
	// lb-abc123
	PublicSlbId *string `json:"PublicSlbId,omitempty" xml:"PublicSlbId,omitempty"`
}

func (s ListGatewayResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayResponseBodyData) GetEnablePublicNet() *bool {
	return s.EnablePublicNet
}

func (s *ListGatewayResponseBodyData) GetFeNodeNumber() *int32 {
	return s.FeNodeNumber
}

func (s *ListGatewayResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewayResponseBodyData) GetGatewayName() *string {
	return s.GatewayName
}

func (s *ListGatewayResponseBodyData) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListGatewayResponseBodyData) GetInternalDomain() *string {
	return s.InternalDomain
}

func (s *ListGatewayResponseBodyData) GetInternalSlbId() *string {
	return s.InternalSlbId
}

func (s *ListGatewayResponseBodyData) GetPrivatezoneId() *string {
	return s.PrivatezoneId
}

func (s *ListGatewayResponseBodyData) GetPublicDomain() *string {
	return s.PublicDomain
}

func (s *ListGatewayResponseBodyData) GetPublicSlbAclId() *string {
	return s.PublicSlbAclId
}

func (s *ListGatewayResponseBodyData) GetPublicSlbId() *string {
	return s.PublicSlbId
}

func (s *ListGatewayResponseBodyData) SetEnablePublicNet(v bool) *ListGatewayResponseBodyData {
	s.EnablePublicNet = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetFeNodeNumber(v int32) *ListGatewayResponseBodyData {
	s.FeNodeNumber = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetGatewayId(v string) *ListGatewayResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetGatewayName(v string) *ListGatewayResponseBodyData {
	s.GatewayName = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetGatewayType(v string) *ListGatewayResponseBodyData {
	s.GatewayType = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetInternalDomain(v string) *ListGatewayResponseBodyData {
	s.InternalDomain = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetInternalSlbId(v string) *ListGatewayResponseBodyData {
	s.InternalSlbId = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetPrivatezoneId(v string) *ListGatewayResponseBodyData {
	s.PrivatezoneId = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetPublicDomain(v string) *ListGatewayResponseBodyData {
	s.PublicDomain = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetPublicSlbAclId(v string) *ListGatewayResponseBodyData {
	s.PublicSlbAclId = &v
	return s
}

func (s *ListGatewayResponseBodyData) SetPublicSlbId(v string) *ListGatewayResponseBodyData {
	s.PublicSlbId = &v
	return s
}

func (s *ListGatewayResponseBodyData) Validate() error {
	return dara.Validate(s)
}
