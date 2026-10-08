// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SaveTaskForSubmittingDomainDeleteRequest
	GetInstanceId() *string
	SetLang(v string) *SaveTaskForSubmittingDomainDeleteRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveTaskForSubmittingDomainDeleteRequest
	GetUserClientIp() *string
}

type SaveTaskForSubmittingDomainDeleteRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// S20181*****85212
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the error message returned by the API. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The user IP address.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForSubmittingDomainDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainDeleteRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainDeleteRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveTaskForSubmittingDomainDeleteRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveTaskForSubmittingDomainDeleteRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveTaskForSubmittingDomainDeleteRequest) SetInstanceId(v string) *SaveTaskForSubmittingDomainDeleteRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteRequest) SetLang(v string) *SaveTaskForSubmittingDomainDeleteRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteRequest) SetUserClientIp(v string) *SaveTaskForSubmittingDomainDeleteRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteRequest) Validate() error {
	return dara.Validate(s)
}
