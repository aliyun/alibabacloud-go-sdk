// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetZoneRecordStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SetZoneRecordStatusRequest
	GetClientToken() *string
	SetLang(v string) *SetZoneRecordStatusRequest
	GetLang() *string
	SetRecordId(v int64) *SetZoneRecordStatusRequest
	GetRecordId() *int64
	SetStatus(v string) *SetZoneRecordStatusRequest
	GetStatus() *string
	SetUserClientIp(v string) *SetZoneRecordStatusRequest
	GetUserClientIp() *string
}

type SetZoneRecordStatusRequest struct {
	// A client token that is used to ensure the idempotence of the request. The client generates this value. The value must be unique among different requests and can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// The default value is en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 207541****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The status of the DNS record. Valid values:
	//
	// - ENABLE: Enables DNS resolution.
	//
	// - DISABLE: Pauses DNS resolution.
	//
	// This parameter is required.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The client IP address.
	//
	// example:
	//
	// 127.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SetZoneRecordStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetZoneRecordStatusRequest) GoString() string {
	return s.String()
}

func (s *SetZoneRecordStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetZoneRecordStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *SetZoneRecordStatusRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *SetZoneRecordStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *SetZoneRecordStatusRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SetZoneRecordStatusRequest) SetClientToken(v string) *SetZoneRecordStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *SetZoneRecordStatusRequest) SetLang(v string) *SetZoneRecordStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetZoneRecordStatusRequest) SetRecordId(v int64) *SetZoneRecordStatusRequest {
	s.RecordId = &v
	return s
}

func (s *SetZoneRecordStatusRequest) SetStatus(v string) *SetZoneRecordStatusRequest {
	s.Status = &v
	return s
}

func (s *SetZoneRecordStatusRequest) SetUserClientIp(v string) *SetZoneRecordStatusRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetZoneRecordStatusRequest) Validate() error {
	return dara.Validate(s)
}
