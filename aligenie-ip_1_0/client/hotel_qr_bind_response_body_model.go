// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelQrBindResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *HotelQrBindResponseBody
	GetMessage() *string
	SetRequestId(v string) *HotelQrBindResponseBody
	GetRequestId() *string
	SetResult(v *HotelQrBindResponseBodyResult) *HotelQrBindResponseBody
	GetResult() *HotelQrBindResponseBodyResult
	SetStatusCode(v int32) *HotelQrBindResponseBody
	GetStatusCode() *int32
}

type HotelQrBindResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73****9-175A-1324-8202-9FAAB*****A
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *HotelQrBindResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s HotelQrBindResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelQrBindResponseBody) GoString() string {
	return s.String()
}

func (s *HotelQrBindResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelQrBindResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelQrBindResponseBody) GetResult() *HotelQrBindResponseBodyResult {
	return s.Result
}

func (s *HotelQrBindResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotelQrBindResponseBody) SetMessage(v string) *HotelQrBindResponseBody {
	s.Message = &v
	return s
}

func (s *HotelQrBindResponseBody) SetRequestId(v string) *HotelQrBindResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelQrBindResponseBody) SetResult(v *HotelQrBindResponseBodyResult) *HotelQrBindResponseBody {
	s.Result = v
	return s
}

func (s *HotelQrBindResponseBody) SetStatusCode(v int32) *HotelQrBindResponseBody {
	s.StatusCode = &v
	return s
}

func (s *HotelQrBindResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HotelQrBindResponseBodyResult struct {
	OpenDeviceInfo *HotelQrBindResponseBodyResultOpenDeviceInfo `json:"OpenDeviceInfo,omitempty" xml:"OpenDeviceInfo,omitempty" type:"Struct"`
	OpenUserInfo   *HotelQrBindResponseBodyResultOpenUserInfo   `json:"OpenUserInfo,omitempty" xml:"OpenUserInfo,omitempty" type:"Struct"`
}

func (s HotelQrBindResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s HotelQrBindResponseBodyResult) GoString() string {
	return s.String()
}

func (s *HotelQrBindResponseBodyResult) GetOpenDeviceInfo() *HotelQrBindResponseBodyResultOpenDeviceInfo {
	return s.OpenDeviceInfo
}

func (s *HotelQrBindResponseBodyResult) GetOpenUserInfo() *HotelQrBindResponseBodyResultOpenUserInfo {
	return s.OpenUserInfo
}

func (s *HotelQrBindResponseBodyResult) SetOpenDeviceInfo(v *HotelQrBindResponseBodyResultOpenDeviceInfo) *HotelQrBindResponseBodyResult {
	s.OpenDeviceInfo = v
	return s
}

func (s *HotelQrBindResponseBodyResult) SetOpenUserInfo(v *HotelQrBindResponseBodyResultOpenUserInfo) *HotelQrBindResponseBodyResult {
	s.OpenUserInfo = v
	return s
}

func (s *HotelQrBindResponseBodyResult) Validate() error {
	if s.OpenDeviceInfo != nil {
		if err := s.OpenDeviceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.OpenUserInfo != nil {
		if err := s.OpenUserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HotelQrBindResponseBodyResultOpenDeviceInfo struct {
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// xxxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// aaaaaaaa
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s HotelQrBindResponseBodyResultOpenDeviceInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelQrBindResponseBodyResultOpenDeviceInfo) GoString() string {
	return s.String()
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) GetId() *string {
	return s.Id
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) GetIdType() *string {
	return s.IdType
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) SetEncodeKey(v string) *HotelQrBindResponseBodyResultOpenDeviceInfo {
	s.EncodeKey = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) SetEncodeType(v string) *HotelQrBindResponseBodyResultOpenDeviceInfo {
	s.EncodeType = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) SetId(v string) *HotelQrBindResponseBodyResultOpenDeviceInfo {
	s.Id = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) SetIdType(v string) *HotelQrBindResponseBodyResultOpenDeviceInfo {
	s.IdType = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) SetOrganizationId(v string) *HotelQrBindResponseBodyResultOpenDeviceInfo {
	s.OrganizationId = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenDeviceInfo) Validate() error {
	return dara.Validate(s)
}

type HotelQrBindResponseBodyResultOpenUserInfo struct {
	// example:
	//
	// 123
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// example:
	//
	// HOTEL
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// example:
	//
	// xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// aaaaaaaa
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s HotelQrBindResponseBodyResultOpenUserInfo) String() string {
	return dara.Prettify(s)
}

func (s HotelQrBindResponseBodyResultOpenUserInfo) GoString() string {
	return s.String()
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) GetId() *string {
	return s.Id
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) SetEncodeKey(v string) *HotelQrBindResponseBodyResultOpenUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) SetEncodeType(v string) *HotelQrBindResponseBodyResultOpenUserInfo {
	s.EncodeType = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) SetId(v string) *HotelQrBindResponseBodyResultOpenUserInfo {
	s.Id = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) SetIdType(v string) *HotelQrBindResponseBodyResultOpenUserInfo {
	s.IdType = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) SetOrganizationId(v string) *HotelQrBindResponseBodyResultOpenUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *HotelQrBindResponseBodyResultOpenUserInfo) Validate() error {
	return dara.Validate(s)
}
