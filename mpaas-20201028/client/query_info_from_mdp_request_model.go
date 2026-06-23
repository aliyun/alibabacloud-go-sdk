// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInfoFromMdpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryInfoFromMdpRequest
	GetAppId() *string
	SetMobile(v string) *QueryInfoFromMdpRequest
	GetMobile() *string
	SetMobileMd5(v string) *QueryInfoFromMdpRequest
	GetMobileMd5() *string
	SetMobileSha256(v string) *QueryInfoFromMdpRequest
	GetMobileSha256() *string
	SetMobileSm3(v string) *QueryInfoFromMdpRequest
	GetMobileSm3() *string
	SetRiskScene(v string) *QueryInfoFromMdpRequest
	GetRiskScene() *string
	SetWorkspaceId(v string) *QueryInfoFromMdpRequest
	GetWorkspaceId() *string
}

type QueryInfoFromMdpRequest struct {
	// This parameter is required.
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Mobile       *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	MobileMd5    *string `json:"MobileMd5,omitempty" xml:"MobileMd5,omitempty"`
	MobileSha256 *string `json:"MobileSha256,omitempty" xml:"MobileSha256,omitempty"`
	MobileSm3    *string `json:"MobileSm3,omitempty" xml:"MobileSm3,omitempty"`
	// This parameter is required.
	RiskScene *string `json:"RiskScene,omitempty" xml:"RiskScene,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryInfoFromMdpRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInfoFromMdpRequest) GoString() string {
	return s.String()
}

func (s *QueryInfoFromMdpRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryInfoFromMdpRequest) GetMobile() *string {
	return s.Mobile
}

func (s *QueryInfoFromMdpRequest) GetMobileMd5() *string {
	return s.MobileMd5
}

func (s *QueryInfoFromMdpRequest) GetMobileSha256() *string {
	return s.MobileSha256
}

func (s *QueryInfoFromMdpRequest) GetMobileSm3() *string {
	return s.MobileSm3
}

func (s *QueryInfoFromMdpRequest) GetRiskScene() *string {
	return s.RiskScene
}

func (s *QueryInfoFromMdpRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryInfoFromMdpRequest) SetAppId(v string) *QueryInfoFromMdpRequest {
	s.AppId = &v
	return s
}

func (s *QueryInfoFromMdpRequest) SetMobile(v string) *QueryInfoFromMdpRequest {
	s.Mobile = &v
	return s
}

func (s *QueryInfoFromMdpRequest) SetMobileMd5(v string) *QueryInfoFromMdpRequest {
	s.MobileMd5 = &v
	return s
}

func (s *QueryInfoFromMdpRequest) SetMobileSha256(v string) *QueryInfoFromMdpRequest {
	s.MobileSha256 = &v
	return s
}

func (s *QueryInfoFromMdpRequest) SetMobileSm3(v string) *QueryInfoFromMdpRequest {
	s.MobileSm3 = &v
	return s
}

func (s *QueryInfoFromMdpRequest) SetRiskScene(v string) *QueryInfoFromMdpRequest {
	s.RiskScene = &v
	return s
}

func (s *QueryInfoFromMdpRequest) SetWorkspaceId(v string) *QueryInfoFromMdpRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryInfoFromMdpRequest) Validate() error {
	return dara.Validate(s)
}
