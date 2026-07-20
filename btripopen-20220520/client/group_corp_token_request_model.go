// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupCorpTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppSecret(v string) *GroupCorpTokenRequest
	GetAppSecret() *string
	SetCorpId(v string) *GroupCorpTokenRequest
	GetCorpId() *string
	SetSubCorpId(v string) *GroupCorpTokenRequest
	GetSubCorpId() *string
}

type GroupCorpTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ZzljczY5dnFjNDAwVlNofiwoWX5ZWCxlcjVTKnVoZS0
	AppSecret *string `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// open12g9sfbmm5i07v10wDzRSK9w00
	CorpId *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// btripuyxmbg3cs286734u_mow6q
	SubCorpId *string `json:"sub_corp_id,omitempty" xml:"sub_corp_id,omitempty"`
}

func (s GroupCorpTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GroupCorpTokenRequest) GoString() string {
	return s.String()
}

func (s *GroupCorpTokenRequest) GetAppSecret() *string {
	return s.AppSecret
}

func (s *GroupCorpTokenRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *GroupCorpTokenRequest) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *GroupCorpTokenRequest) SetAppSecret(v string) *GroupCorpTokenRequest {
	s.AppSecret = &v
	return s
}

func (s *GroupCorpTokenRequest) SetCorpId(v string) *GroupCorpTokenRequest {
	s.CorpId = &v
	return s
}

func (s *GroupCorpTokenRequest) SetSubCorpId(v string) *GroupCorpTokenRequest {
	s.SubCorpId = &v
	return s
}

func (s *GroupCorpTokenRequest) Validate() error {
	return dara.Validate(s)
}
