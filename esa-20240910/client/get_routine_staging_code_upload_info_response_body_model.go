// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineStagingCodeUploadInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersion(v string) *GetRoutineStagingCodeUploadInfoResponseBody
	GetCodeVersion() *string
	SetOssPostConfig(v map[string]interface{}) *GetRoutineStagingCodeUploadInfoResponseBody
	GetOssPostConfig() map[string]interface{}
	SetRequestId(v string) *GetRoutineStagingCodeUploadInfoResponseBody
	GetRequestId() *string
}

type GetRoutineStagingCodeUploadInfoResponseBody struct {
	// The code version number.
	//
	// example:
	//
	// 1710120201067203242
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The configuration information for uploading a JavaScript file to OSS. When uploading to OSS, you must include the OSSAccessKeyId and x-oss-security-token parameters.
	//
	// example:
	//
	// "OssPostConfig": {
	//
	//             "Url": "http://oss_fake_bucket_url",
	//
	//             "OSSAccessKeyId": "xxx",
	//
	//             "XOssSecurityToken": "",
	//
	//             "key": "sp_esa_rwa/xxx/hello/unstable/routine/index.js",
	//
	//             "callback": "xxx=",
	//
	//             "x:codeDescription": "xxx=",
	//
	//             "policy": "xxx=",
	//
	//             "Signature": "xxx="
	//
	//         }
	OssPostConfig map[string]interface{} `json:"OssPostConfig,omitempty" xml:"OssPostConfig,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRoutineStagingCodeUploadInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineStagingCodeUploadInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoutineStagingCodeUploadInfoResponseBody) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *GetRoutineStagingCodeUploadInfoResponseBody) GetOssPostConfig() map[string]interface{} {
	return s.OssPostConfig
}

func (s *GetRoutineStagingCodeUploadInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRoutineStagingCodeUploadInfoResponseBody) SetCodeVersion(v string) *GetRoutineStagingCodeUploadInfoResponseBody {
	s.CodeVersion = &v
	return s
}

func (s *GetRoutineStagingCodeUploadInfoResponseBody) SetOssPostConfig(v map[string]interface{}) *GetRoutineStagingCodeUploadInfoResponseBody {
	s.OssPostConfig = v
	return s
}

func (s *GetRoutineStagingCodeUploadInfoResponseBody) SetRequestId(v string) *GetRoutineStagingCodeUploadInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoutineStagingCodeUploadInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
