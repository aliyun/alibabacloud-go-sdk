// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretTransferRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLastPartnerKey(v string) *CreateSecretTransferRecordRequest
	GetLastPartnerKey() *string
	SetNumberList(v string) *CreateSecretTransferRecordRequest
	GetNumberList() *string
	SetOriginPartnerKey(v string) *CreateSecretTransferRecordRequest
	GetOriginPartnerKey() *string
}

type CreateSecretTransferRecordRequest struct {
	// 迁移后的新号码池key
	//
	// example:
	//
	// FC111111
	LastPartnerKey *string `json:"LastPartnerKey,omitempty" xml:"LastPartnerKey,omitempty"`
	// 指定号码（单一、批量）迁移;一次最多可迁移500个号码，号码间请换行隔开.
	//
	// example:
	//
	// 198xxxxxxxx
	//
	// 138xxxxxxx
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
	// 要迁移的原号码池key
	//
	// example:
	//
	// FC123456
	OriginPartnerKey *string `json:"OriginPartnerKey,omitempty" xml:"OriginPartnerKey,omitempty"`
}

func (s CreateSecretTransferRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretTransferRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretTransferRecordRequest) GetLastPartnerKey() *string {
	return s.LastPartnerKey
}

func (s *CreateSecretTransferRecordRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *CreateSecretTransferRecordRequest) GetOriginPartnerKey() *string {
	return s.OriginPartnerKey
}

func (s *CreateSecretTransferRecordRequest) SetLastPartnerKey(v string) *CreateSecretTransferRecordRequest {
	s.LastPartnerKey = &v
	return s
}

func (s *CreateSecretTransferRecordRequest) SetNumberList(v string) *CreateSecretTransferRecordRequest {
	s.NumberList = &v
	return s
}

func (s *CreateSecretTransferRecordRequest) SetOriginPartnerKey(v string) *CreateSecretTransferRecordRequest {
	s.OriginPartnerKey = &v
	return s
}

func (s *CreateSecretTransferRecordRequest) Validate() error {
	return dara.Validate(s)
}
