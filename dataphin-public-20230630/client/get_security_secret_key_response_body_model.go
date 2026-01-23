// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecuritySecretKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSecuritySecretKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetSecuritySecretKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSecuritySecretKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSecuritySecretKeyResponseBody
	GetRequestId() *string
	SetSecuritySecretKeyInfo(v *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) *GetSecuritySecretKeyResponseBody
	GetSecuritySecretKeyInfo() *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo
	SetSuccess(v bool) *GetSecuritySecretKeyResponseBody
	GetSuccess() *bool
}

type GetSecuritySecretKeyResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId             *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecuritySecretKeyInfo *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo `json:"SecuritySecretKeyInfo,omitempty" xml:"SecuritySecretKeyInfo,omitempty" type:"Struct"`
	Success               *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSecuritySecretKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySecretKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecuritySecretKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSecuritySecretKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSecuritySecretKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSecuritySecretKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSecuritySecretKeyResponseBody) GetSecuritySecretKeyInfo() *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	return s.SecuritySecretKeyInfo
}

func (s *GetSecuritySecretKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSecuritySecretKeyResponseBody) SetCode(v string) *GetSecuritySecretKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBody) SetHttpStatusCode(v int32) *GetSecuritySecretKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBody) SetMessage(v string) *GetSecuritySecretKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBody) SetRequestId(v string) *GetSecuritySecretKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBody) SetSecuritySecretKeyInfo(v *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) *GetSecuritySecretKeyResponseBody {
	s.SecuritySecretKeyInfo = v
	return s
}

func (s *GetSecuritySecretKeyResponseBody) SetSuccess(v bool) *GetSecuritySecretKeyResponseBody {
	s.Success = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBody) Validate() error {
	if s.SecuritySecretKeyInfo != nil {
		if err := s.SecuritySecretKeyInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo struct {
	// example:
	//
	// AES
	AlgorithmType *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
	// example:
	//
	// AES
	AlgorithmTypeAlias *string `json:"AlgorithmTypeAlias,omitempty" xml:"AlgorithmTypeAlias,omitempty"`
	// example:
	//
	// test
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableOpenapiQuery *bool   `json:"EnableOpenapiQuery,omitempty" xml:"EnableOpenapiQuery,omitempty"`
	// example:
	//
	// SYSTEM_GENERATION
	GenerationType *string `json:"GenerationType,omitempty" xml:"GenerationType,omitempty"`
	// example:
	//
	// 1
	Id                *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	IsOwnerManageOnly *bool  `json:"IsOwnerManageOnly,omitempty" xml:"IsOwnerManageOnly,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 张三
	OwnerName     *string   `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	SecretKeyList []*string `json:"SecretKeyList,omitempty" xml:"SecretKeyList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	SubKeyCount *int64 `json:"SubKeyCount,omitempty" xml:"SubKeyCount,omitempty"`
	// example:
	//
	// HASH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GoString() string {
	return s.String()
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetAlgorithmType() *string {
	return s.AlgorithmType
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetAlgorithmTypeAlias() *string {
	return s.AlgorithmTypeAlias
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetDescription() *string {
	return s.Description
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetEnableOpenapiQuery() *bool {
	return s.EnableOpenapiQuery
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetGenerationType() *string {
	return s.GenerationType
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetId() *int64 {
	return s.Id
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetIsOwnerManageOnly() *bool {
	return s.IsOwnerManageOnly
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetName() *string {
	return s.Name
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetOwner() *string {
	return s.Owner
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetSecretKeyList() []*string {
	return s.SecretKeyList
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetSubKeyCount() *int64 {
	return s.SubKeyCount
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) GetType() *string {
	return s.Type
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetAlgorithmType(v string) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.AlgorithmType = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetAlgorithmTypeAlias(v string) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.AlgorithmTypeAlias = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetDescription(v string) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.Description = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetEnableOpenapiQuery(v bool) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.EnableOpenapiQuery = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetGenerationType(v string) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.GenerationType = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetId(v int64) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.Id = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetIsOwnerManageOnly(v bool) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.IsOwnerManageOnly = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetName(v string) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.Name = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetOwner(v string) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.Owner = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetOwnerName(v string) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.OwnerName = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetSecretKeyList(v []*string) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.SecretKeyList = v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetSubKeyCount(v int64) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.SubKeyCount = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) SetType(v string) *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo {
	s.Type = &v
	return s
}

func (s *GetSecuritySecretKeyResponseBodySecuritySecretKeyInfo) Validate() error {
	return dara.Validate(s)
}
