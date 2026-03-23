// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelApThirdAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApAssetId(v int64) *DelApThirdAppRequest
	GetApAssetId() *int64
	SetAppCode(v string) *DelApThirdAppRequest
	GetAppCode() *string
	SetAppName(v string) *DelApThirdAppRequest
	GetAppName() *string
	SetId(v int64) *DelApThirdAppRequest
	GetId() *int64
	SetMac(v string) *DelApThirdAppRequest
	GetMac() *string
}

type DelApThirdAppRequest struct {
	ApAssetId *int64 `json:"ApAssetId,omitempty" xml:"ApAssetId,omitempty"`
	// This parameter is required.
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
}

func (s DelApThirdAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DelApThirdAppRequest) GoString() string {
	return s.String()
}

func (s *DelApThirdAppRequest) GetApAssetId() *int64 {
	return s.ApAssetId
}

func (s *DelApThirdAppRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *DelApThirdAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *DelApThirdAppRequest) GetId() *int64 {
	return s.Id
}

func (s *DelApThirdAppRequest) GetMac() *string {
	return s.Mac
}

func (s *DelApThirdAppRequest) SetApAssetId(v int64) *DelApThirdAppRequest {
	s.ApAssetId = &v
	return s
}

func (s *DelApThirdAppRequest) SetAppCode(v string) *DelApThirdAppRequest {
	s.AppCode = &v
	return s
}

func (s *DelApThirdAppRequest) SetAppName(v string) *DelApThirdAppRequest {
	s.AppName = &v
	return s
}

func (s *DelApThirdAppRequest) SetId(v int64) *DelApThirdAppRequest {
	s.Id = &v
	return s
}

func (s *DelApThirdAppRequest) SetMac(v string) *DelApThirdAppRequest {
	s.Mac = &v
	return s
}

func (s *DelApThirdAppRequest) Validate() error {
	return dara.Validate(s)
}
