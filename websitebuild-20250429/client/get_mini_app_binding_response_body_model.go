// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiniAppBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMiniAppBindingResponseBodyData) *GetMiniAppBindingResponseBody
	GetData() *GetMiniAppBindingResponseBodyData
	SetRequestId(v string) *GetMiniAppBindingResponseBody
	GetRequestId() *string
}

type GetMiniAppBindingResponseBody struct {
	Data *GetMiniAppBindingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMiniAppBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppBindingResponseBody) GoString() string {
	return s.String()
}

func (s *GetMiniAppBindingResponseBody) GetData() *GetMiniAppBindingResponseBodyData {
	return s.Data
}

func (s *GetMiniAppBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMiniAppBindingResponseBody) SetData(v *GetMiniAppBindingResponseBodyData) *GetMiniAppBindingResponseBody {
	s.Data = v
	return s
}

func (s *GetMiniAppBindingResponseBody) SetRequestId(v string) *GetMiniAppBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMiniAppBindingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMiniAppBindingResponseBodyData struct {
	// example:
	//
	// disabled
	AuthStatus *string `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// example:
	//
	// WS20260206134402000001
	BizId               *string            `json:"BizId,omitempty" xml:"BizId,omitempty"`
	IcpFiled            *bool              `json:"IcpFiled,omitempty" xml:"IcpFiled,omitempty"`
	PlatformAppid       *string            `json:"PlatformAppid,omitempty" xml:"PlatformAppid,omitempty"`
	PreviewQrCodeUrl    *string            `json:"PreviewQrCodeUrl,omitempty" xml:"PreviewQrCodeUrl,omitempty"`
	ProductionQrCodeUrl *string            `json:"ProductionQrCodeUrl,omitempty" xml:"ProductionQrCodeUrl,omitempty"`
	SettingValues       map[string]*string `json:"SettingValues,omitempty" xml:"SettingValues,omitempty"`
}

func (s GetMiniAppBindingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppBindingResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMiniAppBindingResponseBodyData) GetAuthStatus() *string {
	return s.AuthStatus
}

func (s *GetMiniAppBindingResponseBodyData) GetBizId() *string {
	return s.BizId
}

func (s *GetMiniAppBindingResponseBodyData) GetIcpFiled() *bool {
	return s.IcpFiled
}

func (s *GetMiniAppBindingResponseBodyData) GetPlatformAppid() *string {
	return s.PlatformAppid
}

func (s *GetMiniAppBindingResponseBodyData) GetPreviewQrCodeUrl() *string {
	return s.PreviewQrCodeUrl
}

func (s *GetMiniAppBindingResponseBodyData) GetProductionQrCodeUrl() *string {
	return s.ProductionQrCodeUrl
}

func (s *GetMiniAppBindingResponseBodyData) GetSettingValues() map[string]*string {
	return s.SettingValues
}

func (s *GetMiniAppBindingResponseBodyData) SetAuthStatus(v string) *GetMiniAppBindingResponseBodyData {
	s.AuthStatus = &v
	return s
}

func (s *GetMiniAppBindingResponseBodyData) SetBizId(v string) *GetMiniAppBindingResponseBodyData {
	s.BizId = &v
	return s
}

func (s *GetMiniAppBindingResponseBodyData) SetIcpFiled(v bool) *GetMiniAppBindingResponseBodyData {
	s.IcpFiled = &v
	return s
}

func (s *GetMiniAppBindingResponseBodyData) SetPlatformAppid(v string) *GetMiniAppBindingResponseBodyData {
	s.PlatformAppid = &v
	return s
}

func (s *GetMiniAppBindingResponseBodyData) SetPreviewQrCodeUrl(v string) *GetMiniAppBindingResponseBodyData {
	s.PreviewQrCodeUrl = &v
	return s
}

func (s *GetMiniAppBindingResponseBodyData) SetProductionQrCodeUrl(v string) *GetMiniAppBindingResponseBodyData {
	s.ProductionQrCodeUrl = &v
	return s
}

func (s *GetMiniAppBindingResponseBodyData) SetSettingValues(v map[string]*string) *GetMiniAppBindingResponseBodyData {
	s.SettingValues = v
	return s
}

func (s *GetMiniAppBindingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
