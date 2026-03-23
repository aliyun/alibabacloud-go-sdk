// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJudgeXingTianBusinessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *JudgeXingTianBusinessRequest
	GetAppCode() *string
	SetAppName(v string) *JudgeXingTianBusinessRequest
	GetAppName() *string
	SetRealmId(v string) *JudgeXingTianBusinessRequest
	GetRealmId() *string
}

type JudgeXingTianBusinessRequest struct {
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	RealmId *string `json:"RealmId,omitempty" xml:"RealmId,omitempty"`
}

func (s JudgeXingTianBusinessRequest) String() string {
	return dara.Prettify(s)
}

func (s JudgeXingTianBusinessRequest) GoString() string {
	return s.String()
}

func (s *JudgeXingTianBusinessRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *JudgeXingTianBusinessRequest) GetAppName() *string {
	return s.AppName
}

func (s *JudgeXingTianBusinessRequest) GetRealmId() *string {
	return s.RealmId
}

func (s *JudgeXingTianBusinessRequest) SetAppCode(v string) *JudgeXingTianBusinessRequest {
	s.AppCode = &v
	return s
}

func (s *JudgeXingTianBusinessRequest) SetAppName(v string) *JudgeXingTianBusinessRequest {
	s.AppName = &v
	return s
}

func (s *JudgeXingTianBusinessRequest) SetRealmId(v string) *JudgeXingTianBusinessRequest {
	s.RealmId = &v
	return s
}

func (s *JudgeXingTianBusinessRequest) Validate() error {
	return dara.Validate(s)
}
