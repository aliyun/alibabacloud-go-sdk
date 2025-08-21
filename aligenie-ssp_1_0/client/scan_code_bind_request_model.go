// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanCodeBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBindReq(v *ScanCodeBindRequestBindReq) *ScanCodeBindRequest
	GetBindReq() *ScanCodeBindRequestBindReq
	SetUserInfo(v *ScanCodeBindRequestUserInfo) *ScanCodeBindRequest
	GetUserInfo() *ScanCodeBindRequestUserInfo
}

type ScanCodeBindRequest struct {
	// This parameter is required.
	BindReq *ScanCodeBindRequestBindReq `json:"BindReq,omitempty" xml:"BindReq,omitempty" type:"Struct"`
	// This parameter is required.
	UserInfo *ScanCodeBindRequestUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ScanCodeBindRequest) String() string {
	return dara.Prettify(s)
}

func (s ScanCodeBindRequest) GoString() string {
	return s.String()
}

func (s *ScanCodeBindRequest) GetBindReq() *ScanCodeBindRequestBindReq {
	return s.BindReq
}

func (s *ScanCodeBindRequest) GetUserInfo() *ScanCodeBindRequestUserInfo {
	return s.UserInfo
}

func (s *ScanCodeBindRequest) SetBindReq(v *ScanCodeBindRequestBindReq) *ScanCodeBindRequest {
	s.BindReq = v
	return s
}

func (s *ScanCodeBindRequest) SetUserInfo(v *ScanCodeBindRequestUserInfo) *ScanCodeBindRequest {
	s.UserInfo = v
	return s
}

func (s *ScanCodeBindRequest) Validate() error {
	return dara.Validate(s)
}

type ScanCodeBindRequestBindReq struct {
	// This parameter is required.
	//
	// example:
	//
	// RnY8v0W0ZVn58ZrUAOr2RD
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// authCode
	//
	// This parameter is required.
	//
	// example:
	//
	// ASdfre
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {"key":"value"}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
}

func (s ScanCodeBindRequestBindReq) String() string {
	return dara.Prettify(s)
}

func (s ScanCodeBindRequestBindReq) GoString() string {
	return s.String()
}

func (s *ScanCodeBindRequestBindReq) GetClientId() *string {
	return s.ClientId
}

func (s *ScanCodeBindRequestBindReq) GetCode() *string {
	return s.Code
}

func (s *ScanCodeBindRequestBindReq) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *ScanCodeBindRequestBindReq) SetClientId(v string) *ScanCodeBindRequestBindReq {
	s.ClientId = &v
	return s
}

func (s *ScanCodeBindRequestBindReq) SetCode(v string) *ScanCodeBindRequestBindReq {
	s.Code = &v
	return s
}

func (s *ScanCodeBindRequestBindReq) SetExtInfo(v string) *ScanCodeBindRequestBindReq {
	s.ExtInfo = &v
	return s
}

func (s *ScanCodeBindRequestBindReq) Validate() error {
	return dara.Validate(s)
}

type ScanCodeBindRequestUserInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// 129****0946
	EncodeKey *string `json:"EncodeKey,omitempty" xml:"EncodeKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PROJECT_ID
	EncodeType *string `json:"EncodeType,omitempty" xml:"EncodeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DAFE****ce3ej=
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OPEN_ID
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// 111
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
}

func (s ScanCodeBindRequestUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ScanCodeBindRequestUserInfo) GoString() string {
	return s.String()
}

func (s *ScanCodeBindRequestUserInfo) GetEncodeKey() *string {
	return s.EncodeKey
}

func (s *ScanCodeBindRequestUserInfo) GetEncodeType() *string {
	return s.EncodeType
}

func (s *ScanCodeBindRequestUserInfo) GetId() *string {
	return s.Id
}

func (s *ScanCodeBindRequestUserInfo) GetIdType() *string {
	return s.IdType
}

func (s *ScanCodeBindRequestUserInfo) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ScanCodeBindRequestUserInfo) SetEncodeKey(v string) *ScanCodeBindRequestUserInfo {
	s.EncodeKey = &v
	return s
}

func (s *ScanCodeBindRequestUserInfo) SetEncodeType(v string) *ScanCodeBindRequestUserInfo {
	s.EncodeType = &v
	return s
}

func (s *ScanCodeBindRequestUserInfo) SetId(v string) *ScanCodeBindRequestUserInfo {
	s.Id = &v
	return s
}

func (s *ScanCodeBindRequestUserInfo) SetIdType(v string) *ScanCodeBindRequestUserInfo {
	s.IdType = &v
	return s
}

func (s *ScanCodeBindRequestUserInfo) SetOrganizationId(v string) *ScanCodeBindRequestUserInfo {
	s.OrganizationId = &v
	return s
}

func (s *ScanCodeBindRequestUserInfo) Validate() error {
	return dara.Validate(s)
}
