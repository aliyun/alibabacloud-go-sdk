// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPluginConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppPluginConfigRequest
	GetBizId() *string
	SetPluginId(v string) *GetAppPluginConfigRequest
	GetPluginId() *string
}

type GetAppPluginConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1bae9ceaceea432d91c7069fab0dfc02
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
}

func (s GetAppPluginConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppPluginConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAppPluginConfigRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppPluginConfigRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *GetAppPluginConfigRequest) SetBizId(v string) *GetAppPluginConfigRequest {
	s.BizId = &v
	return s
}

func (s *GetAppPluginConfigRequest) SetPluginId(v string) *GetAppPluginConfigRequest {
	s.PluginId = &v
	return s
}

func (s *GetAppPluginConfigRequest) Validate() error {
	return dara.Validate(s)
}
