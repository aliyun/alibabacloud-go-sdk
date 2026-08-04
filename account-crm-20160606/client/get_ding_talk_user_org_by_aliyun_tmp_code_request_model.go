// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingTalkUserOrgByAliyunTmpCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetDingTalkUserOrgByAliyunTmpCodeRequest
	GetAppName() *string
	SetDingTalkChannel(v string) *GetDingTalkUserOrgByAliyunTmpCodeRequest
	GetDingTalkChannel() *string
	SetTmpCode(v string) *GetDingTalkUserOrgByAliyunTmpCodeRequest
	GetTmpCode() *string
	SetVersion(v string) *GetDingTalkUserOrgByAliyunTmpCodeRequest
	GetVersion() *string
}

type GetDingTalkUserOrgByAliyunTmpCodeRequest struct {
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DingTalkChannel *string `json:"DingTalkChannel,omitempty" xml:"DingTalkChannel,omitempty"`
	TmpCode         *string `json:"TmpCode,omitempty" xml:"TmpCode,omitempty"`
	Version         *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetDingTalkUserOrgByAliyunTmpCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingTalkUserOrgByAliyunTmpCodeRequest) GoString() string {
	return s.String()
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeRequest) GetDingTalkChannel() *string {
	return s.DingTalkChannel
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeRequest) GetTmpCode() *string {
	return s.TmpCode
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeRequest) GetVersion() *string {
	return s.Version
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeRequest) SetAppName(v string) *GetDingTalkUserOrgByAliyunTmpCodeRequest {
	s.AppName = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeRequest) SetDingTalkChannel(v string) *GetDingTalkUserOrgByAliyunTmpCodeRequest {
	s.DingTalkChannel = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeRequest) SetTmpCode(v string) *GetDingTalkUserOrgByAliyunTmpCodeRequest {
	s.TmpCode = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeRequest) SetVersion(v string) *GetDingTalkUserOrgByAliyunTmpCodeRequest {
	s.Version = &v
	return s
}

func (s *GetDingTalkUserOrgByAliyunTmpCodeRequest) Validate() error {
	return dara.Validate(s)
}
