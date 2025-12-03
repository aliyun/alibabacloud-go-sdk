// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadFilePolicyForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowRetry(v bool) *GenerateUploadFilePolicyForPartnerResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GenerateUploadFilePolicyForPartnerResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GenerateUploadFilePolicyForPartnerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GenerateUploadFilePolicyForPartnerResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GenerateUploadFilePolicyForPartnerResponseBody
	GetErrorArgs() []interface{}
	SetErrorCode(v string) *GenerateUploadFilePolicyForPartnerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GenerateUploadFilePolicyForPartnerResponseBody
	GetErrorMsg() *string
	SetHttpStatusCode(v int32) *GenerateUploadFilePolicyForPartnerResponseBody
	GetHttpStatusCode() *int32
	SetModule(v *GenerateUploadFilePolicyForPartnerResponseBodyModule) *GenerateUploadFilePolicyForPartnerResponseBody
	GetModule() *GenerateUploadFilePolicyForPartnerResponseBodyModule
	SetRequestId(v string) *GenerateUploadFilePolicyForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateUploadFilePolicyForPartnerResponseBody
	GetSuccess() *bool
}

type GenerateUploadFilePolicyForPartnerResponseBody struct {
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// live
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string       `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// example:
	//
	// 500
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// User not authorized to operate on the specified resource.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Module         *GenerateUploadFilePolicyForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// 6254E13A-A79F-5786-9C75-7590727342C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateUploadFilePolicyForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadFilePolicyForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) GetModule() *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	return s.Module
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetAllowRetry(v bool) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetAppName(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.AppName = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetDynamicCode(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetDynamicMessage(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetErrorArgs(v []interface{}) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetErrorCode(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetErrorMsg(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetHttpStatusCode(v int32) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetModule(v *GenerateUploadFilePolicyForPartnerResponseBodyModule) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetRequestId(v string) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) SetSuccess(v bool) *GenerateUploadFilePolicyForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateUploadFilePolicyForPartnerResponseBodyModule struct {
	// example:
	//
	// LTAI5tQPEXqDVu7794Bvw2xM
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// XXXXXXX
	EncodedPolicy *string `json:"EncodedPolicy,omitempty" xml:"EncodedPolicy,omitempty"`
	// example:
	//
	// 1719112842
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// website_partner_leads/website/xxxxxx/xxxxxx
	FileDir *string `json:"FileDir,omitempty" xml:"FileDir,omitempty"`
	// example:
	//
	// //xx-xxx-partner.oss-cn-zhangjiakou.aliyuncs.com/
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// https://msea-website-partner.oss-cn-zhangjiakou.aliyuncs.com/website_xxxx_party_leads/website/xxxx/xxxx/2024/06/25/website_partner_third_party_leads_01?Expires=1719868413&OSSAccessKeyId=LTAI5tAnyDDDDD&Signature=XXXX
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	// example:
	//
	// qQb34p8lIXcSFtog2y0H08bC0OI=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GenerateUploadFilePolicyForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadFilePolicyForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) GetAccessId() *string {
	return s.AccessId
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) GetEncodedPolicy() *string {
	return s.EncodedPolicy
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) GetFileDir() *string {
	return s.FileDir
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) GetHost() *string {
	return s.Host
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) GetOssUrl() *string {
	return s.OssUrl
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) GetSignature() *string {
	return s.Signature
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetAccessId(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.AccessId = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetEncodedPolicy(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.EncodedPolicy = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetExpireTime(v int64) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.ExpireTime = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetFileDir(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.FileDir = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetHost(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.Host = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetOssUrl(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.OssUrl = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) SetSignature(v string) *GenerateUploadFilePolicyForPartnerResponseBodyModule {
	s.Signature = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
