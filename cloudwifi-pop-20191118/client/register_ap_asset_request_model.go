// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterApAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApGroupUUId(v string) *RegisterApAssetRequest
	GetApGroupUUId() *string
	SetAppCode(v string) *RegisterApAssetRequest
	GetAppCode() *string
	SetAppName(v string) *RegisterApAssetRequest
	GetAppName() *string
	SetId(v int64) *RegisterApAssetRequest
	GetId() *int64
	SetMac(v string) *RegisterApAssetRequest
	GetMac() *string
	SetSerialNo(v string) *RegisterApAssetRequest
	GetSerialNo() *string
}

type RegisterApAssetRequest struct {
	// This parameter is required.
	ApGroupUUId *string `json:"ApGroupUUId,omitempty" xml:"ApGroupUUId,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// This parameter is required.
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
}

func (s RegisterApAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterApAssetRequest) GoString() string {
	return s.String()
}

func (s *RegisterApAssetRequest) GetApGroupUUId() *string {
	return s.ApGroupUUId
}

func (s *RegisterApAssetRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *RegisterApAssetRequest) GetAppName() *string {
	return s.AppName
}

func (s *RegisterApAssetRequest) GetId() *int64 {
	return s.Id
}

func (s *RegisterApAssetRequest) GetMac() *string {
	return s.Mac
}

func (s *RegisterApAssetRequest) GetSerialNo() *string {
	return s.SerialNo
}

func (s *RegisterApAssetRequest) SetApGroupUUId(v string) *RegisterApAssetRequest {
	s.ApGroupUUId = &v
	return s
}

func (s *RegisterApAssetRequest) SetAppCode(v string) *RegisterApAssetRequest {
	s.AppCode = &v
	return s
}

func (s *RegisterApAssetRequest) SetAppName(v string) *RegisterApAssetRequest {
	s.AppName = &v
	return s
}

func (s *RegisterApAssetRequest) SetId(v int64) *RegisterApAssetRequest {
	s.Id = &v
	return s
}

func (s *RegisterApAssetRequest) SetMac(v string) *RegisterApAssetRequest {
	s.Mac = &v
	return s
}

func (s *RegisterApAssetRequest) SetSerialNo(v string) *RegisterApAssetRequest {
	s.SerialNo = &v
	return s
}

func (s *RegisterApAssetRequest) Validate() error {
	return dara.Validate(s)
}
