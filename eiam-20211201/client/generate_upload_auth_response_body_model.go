// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *GenerateUploadAuthResponseBody
	GetAccessId() *string
	SetDownLoadUrl(v string) *GenerateUploadAuthResponseBody
	GetDownLoadUrl() *string
	SetEncryptedKey(v string) *GenerateUploadAuthResponseBody
	GetEncryptedKey() *string
	SetExpire(v int64) *GenerateUploadAuthResponseBody
	GetExpire() *int64
	SetHost(v string) *GenerateUploadAuthResponseBody
	GetHost() *string
	SetKey(v string) *GenerateUploadAuthResponseBody
	GetKey() *string
	SetPlaintextKey(v string) *GenerateUploadAuthResponseBody
	GetPlaintextKey() *string
	SetPolicy(v string) *GenerateUploadAuthResponseBody
	GetPolicy() *string
	SetRequestId(v string) *GenerateUploadAuthResponseBody
	GetRequestId() *string
	SetSecurityToken(v string) *GenerateUploadAuthResponseBody
	GetSecurityToken() *string
	SetSignature(v string) *GenerateUploadAuthResponseBody
	GetSignature() *string
}

type GenerateUploadAuthResponseBody struct {
	// 认证的AccessId
	//
	// example:
	//
	// STS.NYgAmE3cyPoMDxtWgtwG3xxxx
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// 预下载地址
	//
	// example:
	//
	// https://temp.oss.aliyuncs.com/idaas_ly77wa2oexrciw5v672vizxxxx/file_import/68866d21-0ab7-450d-b9e6-5b1eafe06xxxx
	DownLoadUrl *string `json:"DownLoadUrl,omitempty" xml:"DownLoadUrl,omitempty"`
	// example:
	//
	// eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIiwia2lkIjoia2V5X3Z1ZWhjbmh2NWppcGhmZGJwcWpqd3dsemFhIn0..YetpxsPdqdLvAy6G.0Zy5meoTzvCuNfA_0w7E9ItY2uGu1v1BxR9i98KeHXv_P-sm9w1q0XPf5Fw.55V_jFq2t2ZHdjg5c7uxxxx
	EncryptedKey *string `json:"EncryptedKey,omitempty" xml:"EncryptedKey,omitempty"`
	// 过期时间
	//
	// example:
	//
	// 1766470716
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// bucket地址host
	//
	// example:
	//
	// https://temp.oss.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// 认证对应的key
	//
	// example:
	//
	// idaas-csv://idaas_ly77wa2oexrciw5v672vizxxxx
	//
	// /file_import/68866d21-0ab7-450d-b9e6-5b1eafxxxx"
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// fBLqWEvq3SbCj1cX/rtZoSPDVduOWdloOO2VN2+5Sxxxx
	PlaintextKey *string `json:"PlaintextKey,omitempty" xml:"PlaintextKey,omitempty"`
	// 认证的policy
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNS0xMi0yM1QwNjoxODozNi4zNTZaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjBdLHsiYnVja2V0IjoidGVtcC1pZGFhcy1laWFtLWNuLWhhbmd6aG91In0seyJrZXkiOiJpZGFhc19seTc3d2Eyb2V4cmNpdzV2Njcydml6eG12ZS9maWxlX2ltcG9ydC82ODg2NmQyMS0wYWI3LTQ1MGQtYjllNi01YjFlYWZlMDYzZTEifSx7Ingtb3NzLWZvcmJpZC1vdmVyd3JpdGUiOiJ0cnVlIn0seyJ4LW9zcy1vYmplY3QtYWNsIjoicHJpdmF0ZSJxxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// CAISzgJ1q6Ft5B2yfSjIr5rSCtfx3rxY562mRl7Fs2che8gfpbLg1zz2IHhMfXVpA+Afv/sxlG5Q7/wdlrp6SJtIXleCZtF94oxN9h2gb4fb4wgFPgjY08/LI3OaLjKm9u2wCryLYbGwU/OpbE++5U0X6LDmdDKkckW4OJmS8/BOZcgWWQ/KBlgvRq0hRG1YpdQdKGHaONu0LxfumRCwNkdzvRdmgm4NgsbWgO/ks0aO0wehm7BO+N6gfcD9NvMBZskvD42Hu8VtbbfE3SJq7BxHybx7lqQs+02c5onEXwALs0zXbLSErIU2dlBjH68hAOtFquPgnPtzt6nJkID62421pmiSr561rumAtyikcIvBXr5RHT3LoP1LA5UhHC1UotFVgGOaCXLbtuArXptaY/JiNL/0hFEpVt7knInNpUbntINy5f5fqzNMlShqKOXTK93xGoABgfahfEFC23BhLp3NnBxnaG/psewhUfRg/wYS9oE268EST0qPq0ZvzmZjsmtbtnFL8takNDOIdutBZeb9nngkEi7tYyVcwoYBnbJ3sopnFEVozB2VO5XuRHLBkOfB+2z1zD91KtoDajJxpT295Qm0ndFALS1tCGI452yRIrCHynggxxxxx
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// 认证的签名
	//
	// example:
	//
	// t3JyykEKg3kHQuUrhaXYxtboUxxxxx
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GenerateUploadAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadAuthResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadAuthResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *GenerateUploadAuthResponseBody) GetDownLoadUrl() *string {
	return s.DownLoadUrl
}

func (s *GenerateUploadAuthResponseBody) GetEncryptedKey() *string {
	return s.EncryptedKey
}

func (s *GenerateUploadAuthResponseBody) GetExpire() *int64 {
	return s.Expire
}

func (s *GenerateUploadAuthResponseBody) GetHost() *string {
	return s.Host
}

func (s *GenerateUploadAuthResponseBody) GetKey() *string {
	return s.Key
}

func (s *GenerateUploadAuthResponseBody) GetPlaintextKey() *string {
	return s.PlaintextKey
}

func (s *GenerateUploadAuthResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *GenerateUploadAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateUploadAuthResponseBody) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GenerateUploadAuthResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *GenerateUploadAuthResponseBody) SetAccessId(v string) *GenerateUploadAuthResponseBody {
	s.AccessId = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) SetDownLoadUrl(v string) *GenerateUploadAuthResponseBody {
	s.DownLoadUrl = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) SetEncryptedKey(v string) *GenerateUploadAuthResponseBody {
	s.EncryptedKey = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) SetExpire(v int64) *GenerateUploadAuthResponseBody {
	s.Expire = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) SetHost(v string) *GenerateUploadAuthResponseBody {
	s.Host = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) SetKey(v string) *GenerateUploadAuthResponseBody {
	s.Key = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) SetPlaintextKey(v string) *GenerateUploadAuthResponseBody {
	s.PlaintextKey = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) SetPolicy(v string) *GenerateUploadAuthResponseBody {
	s.Policy = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) SetRequestId(v string) *GenerateUploadAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) SetSecurityToken(v string) *GenerateUploadAuthResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) SetSignature(v string) *GenerateUploadAuthResponseBody {
	s.Signature = &v
	return s
}

func (s *GenerateUploadAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
