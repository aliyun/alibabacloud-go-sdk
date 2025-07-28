// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteZoneRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteZoneRecordRequest
	GetClientToken() *string
	SetLang(v string) *DeleteZoneRecordRequest
	GetLang() *string
	SetRecordId(v int64) *DeleteZoneRecordRequest
	GetRecordId() *int64
	SetUserClientIp(v string) *DeleteZoneRecordRequest
	GetUserClientIp() *string
}

type DeleteZoneRecordRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
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
	// 306279****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteZoneRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteZoneRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteZoneRecordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteZoneRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteZoneRecordRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DeleteZoneRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DeleteZoneRecordRequest) SetClientToken(v string) *DeleteZoneRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteZoneRecordRequest) SetLang(v string) *DeleteZoneRecordRequest {
	s.Lang = &v
	return s
}

func (s *DeleteZoneRecordRequest) SetRecordId(v int64) *DeleteZoneRecordRequest {
	s.RecordId = &v
	return s
}

func (s *DeleteZoneRecordRequest) SetUserClientIp(v string) *DeleteZoneRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteZoneRecordRequest) Validate() error {
	return dara.Validate(s)
}
