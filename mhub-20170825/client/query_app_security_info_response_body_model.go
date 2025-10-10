// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppSecurityInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppSecurityInfo(v *QueryAppSecurityInfoResponseBodyAppSecurityInfo) *QueryAppSecurityInfoResponseBody
	GetAppSecurityInfo() *QueryAppSecurityInfoResponseBodyAppSecurityInfo
	SetRequestId(v string) *QueryAppSecurityInfoResponseBody
	GetRequestId() *string
}

type QueryAppSecurityInfoResponseBody struct {
	AppSecurityInfo *QueryAppSecurityInfoResponseBodyAppSecurityInfo `json:"AppSecurityInfo,omitempty" xml:"AppSecurityInfo,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAppSecurityInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAppSecurityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppSecurityInfoResponseBody) GetAppSecurityInfo() *QueryAppSecurityInfoResponseBodyAppSecurityInfo {
	return s.AppSecurityInfo
}

func (s *QueryAppSecurityInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAppSecurityInfoResponseBody) SetAppSecurityInfo(v *QueryAppSecurityInfoResponseBodyAppSecurityInfo) *QueryAppSecurityInfoResponseBody {
	s.AppSecurityInfo = v
	return s
}

func (s *QueryAppSecurityInfoResponseBody) SetRequestId(v string) *QueryAppSecurityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAppSecurityInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryAppSecurityInfoResponseBodyAppSecurityInfo struct {
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// abc123abc123
	AppSecret *string                                                   `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	ExtConfig *QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig `json:"ExtConfig,omitempty" xml:"ExtConfig,omitempty" type:"Struct"`
}

func (s QueryAppSecurityInfoResponseBodyAppSecurityInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryAppSecurityInfoResponseBodyAppSecurityInfo) GoString() string {
	return s.String()
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfo) GetAppKey() *string {
	return s.AppKey
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfo) GetAppSecret() *string {
	return s.AppSecret
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfo) GetExtConfig() *QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig {
	return s.ExtConfig
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfo) SetAppKey(v string) *QueryAppSecurityInfoResponseBodyAppSecurityInfo {
	s.AppKey = &v
	return s
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfo) SetAppSecret(v string) *QueryAppSecurityInfoResponseBodyAppSecurityInfo {
	s.AppSecret = &v
	return s
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfo) SetExtConfig(v *QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig) *QueryAppSecurityInfoResponseBodyAppSecurityInfo {
	s.ExtConfig = v
	return s
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfo) Validate() error {
	return dara.Validate(s)
}

type QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig struct {
	TlogRsaSecret *string `json:"TlogRsaSecret,omitempty" xml:"TlogRsaSecret,omitempty"`
}

func (s QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig) String() string {
	return dara.Prettify(s)
}

func (s QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig) GoString() string {
	return s.String()
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig) GetTlogRsaSecret() *string {
	return s.TlogRsaSecret
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig) SetTlogRsaSecret(v string) *QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig {
	s.TlogRsaSecret = &v
	return s
}

func (s *QueryAppSecurityInfoResponseBodyAppSecurityInfoExtConfig) Validate() error {
	return dara.Validate(s)
}
