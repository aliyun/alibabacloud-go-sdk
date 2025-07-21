// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantDeviceOtaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTenantDeviceOtaInfoResponseBody
	GetCode() *string
	SetData(v *ListTenantDeviceOtaInfoResponseBodyData) *ListTenantDeviceOtaInfoResponseBody
	GetData() *ListTenantDeviceOtaInfoResponseBodyData
	SetMessage(v string) *ListTenantDeviceOtaInfoResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListTenantDeviceOtaInfoResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTenantDeviceOtaInfoResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTenantDeviceOtaInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListTenantDeviceOtaInfoResponseBody
	GetTotalCount() *int64
}

type ListTenantDeviceOtaInfoResponseBody struct {
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *ListTenantDeviceOtaInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message    *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTenantDeviceOtaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTenantDeviceOtaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListTenantDeviceOtaInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTenantDeviceOtaInfoResponseBody) GetData() *ListTenantDeviceOtaInfoResponseBodyData {
	return s.Data
}

func (s *ListTenantDeviceOtaInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTenantDeviceOtaInfoResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTenantDeviceOtaInfoResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTenantDeviceOtaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTenantDeviceOtaInfoResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetCode(v string) *ListTenantDeviceOtaInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetData(v *ListTenantDeviceOtaInfoResponseBodyData) *ListTenantDeviceOtaInfoResponseBody {
	s.Data = v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetMessage(v string) *ListTenantDeviceOtaInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetPageNumber(v int32) *ListTenantDeviceOtaInfoResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetPageSize(v int32) *ListTenantDeviceOtaInfoResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetRequestId(v string) *ListTenantDeviceOtaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) SetTotalCount(v int64) *ListTenantDeviceOtaInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTenantDeviceOtaInfoResponseBodyData struct {
	TenantDeviceOtaInfos []*ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos `json:"TenantDeviceOtaInfos,omitempty" xml:"TenantDeviceOtaInfos,omitempty" type:"Repeated"`
}

func (s ListTenantDeviceOtaInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTenantDeviceOtaInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTenantDeviceOtaInfoResponseBodyData) GetTenantDeviceOtaInfos() []*ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos {
	return s.TenantDeviceOtaInfos
}

func (s *ListTenantDeviceOtaInfoResponseBodyData) SetTenantDeviceOtaInfos(v []*ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) *ListTenantDeviceOtaInfoResponseBodyData {
	s.TenantDeviceOtaInfos = v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos struct {
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	DeviceId       *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Model          *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) String() string {
	return dara.Prettify(s)
}

func (s ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) GoString() string {
	return s.String()
}

func (s *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) GetDeviceId() *string {
	return s.DeviceId
}

func (s *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) GetModel() *string {
	return s.Model
}

func (s *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) SetCurrentVersion(v string) *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos {
	s.CurrentVersion = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) SetDeviceId(v string) *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos {
	s.DeviceId = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) SetModel(v string) *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos {
	s.Model = &v
	return s
}

func (s *ListTenantDeviceOtaInfoResponseBodyDataTenantDeviceOtaInfos) Validate() error {
	return dara.Validate(s)
}
