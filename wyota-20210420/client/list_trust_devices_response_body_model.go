// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrustDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTrustDevicesResponseBody
	GetCode() *string
	SetData(v []*ListTrustDevicesResponseBodyData) *ListTrustDevicesResponseBody
	GetData() []*ListTrustDevicesResponseBodyData
	SetMessage(v string) *ListTrustDevicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTrustDevicesResponseBody
	GetRequestId() *string
}

type ListTrustDevicesResponseBody struct {
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*ListTrustDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTrustDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTrustDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTrustDevicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTrustDevicesResponseBody) GetData() []*ListTrustDevicesResponseBodyData {
	return s.Data
}

func (s *ListTrustDevicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTrustDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTrustDevicesResponseBody) SetCode(v string) *ListTrustDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTrustDevicesResponseBody) SetData(v []*ListTrustDevicesResponseBodyData) *ListTrustDevicesResponseBody {
	s.Data = v
	return s
}

func (s *ListTrustDevicesResponseBody) SetMessage(v string) *ListTrustDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTrustDevicesResponseBody) SetRequestId(v string) *ListTrustDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTrustDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTrustDevicesResponseBodyData struct {
	Model    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Uuid     *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListTrustDevicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTrustDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTrustDevicesResponseBodyData) GetModel() *string {
	return s.Model
}

func (s *ListTrustDevicesResponseBodyData) GetSerialNo() *string {
	return s.SerialNo
}

func (s *ListTrustDevicesResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListTrustDevicesResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *ListTrustDevicesResponseBodyData) SetModel(v string) *ListTrustDevicesResponseBodyData {
	s.Model = &v
	return s
}

func (s *ListTrustDevicesResponseBodyData) SetSerialNo(v string) *ListTrustDevicesResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *ListTrustDevicesResponseBodyData) SetTenantId(v string) *ListTrustDevicesResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *ListTrustDevicesResponseBodyData) SetUuid(v string) *ListTrustDevicesResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListTrustDevicesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
