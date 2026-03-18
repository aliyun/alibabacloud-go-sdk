// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrgInitResultDTO interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *OrgInitResultDTO
	GetApiKey() *string
	SetClientId(v int64) *OrgInitResultDTO
	GetClientId() *int64
	SetClientUuid(v string) *OrgInitResultDTO
	GetClientUuid() *string
	SetIsNew(v bool) *OrgInitResultDTO
	GetIsNew() *bool
	SetOrgId(v int64) *OrgInitResultDTO
	GetOrgId() *int64
	SetUserId(v int64) *OrgInitResultDTO
	GetUserId() *int64
}

type OrgInitResultDTO struct {
	// example:
	//
	// sk-xxxxxxxxxxxxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// example:
	//
	// 1
	ClientId *int64 `json:"clientId,omitempty" xml:"clientId,omitempty"`
	// example:
	//
	// st-xxxx
	ClientUuid *string `json:"clientUuid,omitempty" xml:"clientUuid,omitempty"`
	// example:
	//
	// true
	IsNew *bool `json:"isNew,omitempty" xml:"isNew,omitempty"`
	// example:
	//
	// 1
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// example:
	//
	// 1
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s OrgInitResultDTO) String() string {
	return dara.Prettify(s)
}

func (s OrgInitResultDTO) GoString() string {
	return s.String()
}

func (s *OrgInitResultDTO) GetApiKey() *string {
	return s.ApiKey
}

func (s *OrgInitResultDTO) GetClientId() *int64 {
	return s.ClientId
}

func (s *OrgInitResultDTO) GetClientUuid() *string {
	return s.ClientUuid
}

func (s *OrgInitResultDTO) GetIsNew() *bool {
	return s.IsNew
}

func (s *OrgInitResultDTO) GetOrgId() *int64 {
	return s.OrgId
}

func (s *OrgInitResultDTO) GetUserId() *int64 {
	return s.UserId
}

func (s *OrgInitResultDTO) SetApiKey(v string) *OrgInitResultDTO {
	s.ApiKey = &v
	return s
}

func (s *OrgInitResultDTO) SetClientId(v int64) *OrgInitResultDTO {
	s.ClientId = &v
	return s
}

func (s *OrgInitResultDTO) SetClientUuid(v string) *OrgInitResultDTO {
	s.ClientUuid = &v
	return s
}

func (s *OrgInitResultDTO) SetIsNew(v bool) *OrgInitResultDTO {
	s.IsNew = &v
	return s
}

func (s *OrgInitResultDTO) SetOrgId(v int64) *OrgInitResultDTO {
	s.OrgId = &v
	return s
}

func (s *OrgInitResultDTO) SetUserId(v int64) *OrgInitResultDTO {
	s.UserId = &v
	return s
}

func (s *OrgInitResultDTO) Validate() error {
	return dara.Validate(s)
}
