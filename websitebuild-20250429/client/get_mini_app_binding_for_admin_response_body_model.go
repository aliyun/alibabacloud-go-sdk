// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiniAppBindingForAdminResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMiniAppBindingForAdminResponseBodyData) *GetMiniAppBindingForAdminResponseBody
	GetData() *GetMiniAppBindingForAdminResponseBodyData
	SetRequestId(v string) *GetMiniAppBindingForAdminResponseBody
	GetRequestId() *string
}

type GetMiniAppBindingForAdminResponseBody struct {
	// Request result.
	Data *GetMiniAppBindingForAdminResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMiniAppBindingForAdminResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppBindingForAdminResponseBody) GoString() string {
	return s.String()
}

func (s *GetMiniAppBindingForAdminResponseBody) GetData() *GetMiniAppBindingForAdminResponseBodyData {
	return s.Data
}

func (s *GetMiniAppBindingForAdminResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMiniAppBindingForAdminResponseBody) SetData(v *GetMiniAppBindingForAdminResponseBodyData) *GetMiniAppBindingForAdminResponseBody {
	s.Data = v
	return s
}

func (s *GetMiniAppBindingForAdminResponseBody) SetRequestId(v string) *GetMiniAppBindingForAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMiniAppBindingForAdminResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMiniAppBindingForAdminResponseBodyData struct {
	// Authorization status
	//
	// example:
	//
	// AUTHORIZED
	AuthStatus *string `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// Business ID
	//
	// example:
	//
	// WS20250801004817000002
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Miniapp ID
	//
	// example:
	//
	// xxxx
	PlatformAppid *string `json:"PlatformAppid,omitempty" xml:"PlatformAppid,omitempty"`
}

func (s GetMiniAppBindingForAdminResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppBindingForAdminResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMiniAppBindingForAdminResponseBodyData) GetAuthStatus() *string {
	return s.AuthStatus
}

func (s *GetMiniAppBindingForAdminResponseBodyData) GetBizId() *string {
	return s.BizId
}

func (s *GetMiniAppBindingForAdminResponseBodyData) GetPlatformAppid() *string {
	return s.PlatformAppid
}

func (s *GetMiniAppBindingForAdminResponseBodyData) SetAuthStatus(v string) *GetMiniAppBindingForAdminResponseBodyData {
	s.AuthStatus = &v
	return s
}

func (s *GetMiniAppBindingForAdminResponseBodyData) SetBizId(v string) *GetMiniAppBindingForAdminResponseBodyData {
	s.BizId = &v
	return s
}

func (s *GetMiniAppBindingForAdminResponseBodyData) SetPlatformAppid(v string) *GetMiniAppBindingForAdminResponseBodyData {
	s.PlatformAppid = &v
	return s
}

func (s *GetMiniAppBindingForAdminResponseBodyData) Validate() error {
	return dara.Validate(s)
}
