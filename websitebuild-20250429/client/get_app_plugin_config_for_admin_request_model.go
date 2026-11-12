// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPluginConfigForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppPluginConfigForAdminRequest
	GetBizId() *string
	SetPluginId(v string) *GetAppPluginConfigForAdminRequest
	GetPluginId() *string
}

type GetAppPluginConfigForAdminRequest struct {
	// The business ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// WD20250821161210000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The gateway plugin ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
}

func (s GetAppPluginConfigForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginConfigForAdminRequest) GoString() string {
	return s.String()
}

func (s *GetAppPluginConfigForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppPluginConfigForAdminRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *GetAppPluginConfigForAdminRequest) SetBizId(v string) *GetAppPluginConfigForAdminRequest {
	s.BizId = &v
	return s
}

func (s *GetAppPluginConfigForAdminRequest) SetPluginId(v string) *GetAppPluginConfigForAdminRequest {
	s.PluginId = &v
	return s
}

func (s *GetAppPluginConfigForAdminRequest) Validate() error {
	return dara.Validate(s)
}
