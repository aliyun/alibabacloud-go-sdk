// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v *DescribeUserConfigsResponseBodyConfigs) *DescribeUserConfigsResponseBody
	GetConfigs() *DescribeUserConfigsResponseBodyConfigs
	SetRequestId(v string) *DescribeUserConfigsResponseBody
	GetRequestId() *string
}

type DescribeUserConfigsResponseBody struct {
	// The configurations of the specified feature.
	Configs *DescribeUserConfigsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9BCC7BAA-ACBE-45E5-83F0-98BF7E693E84
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsResponseBody) GetConfigs() *DescribeUserConfigsResponseBodyConfigs {
	return s.Configs
}

func (s *DescribeUserConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserConfigsResponseBody) SetConfigs(v *DescribeUserConfigsResponseBodyConfigs) *DescribeUserConfigsResponseBody {
	s.Configs = v
	return s
}

func (s *DescribeUserConfigsResponseBody) SetRequestId(v string) *DescribeUserConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserConfigsResponseBodyConfigs struct {
	// The configurations of Object Storage Service (OSS).
	OssLogConfig *DescribeUserConfigsResponseBodyConfigsOssLogConfig `json:"OssLogConfig,omitempty" xml:"OssLogConfig,omitempty" type:"Struct"`
	// The configurations of Web Application Firewall (WAF).
	WafConfig *DescribeUserConfigsResponseBodyConfigsWafConfig `json:"WafConfig,omitempty" xml:"WafConfig,omitempty" type:"Struct"`
}

func (s DescribeUserConfigsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConfigsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsResponseBodyConfigs) GetOssLogConfig() *DescribeUserConfigsResponseBodyConfigsOssLogConfig {
	return s.OssLogConfig
}

func (s *DescribeUserConfigsResponseBodyConfigs) GetWafConfig() *DescribeUserConfigsResponseBodyConfigsWafConfig {
	return s.WafConfig
}

func (s *DescribeUserConfigsResponseBodyConfigs) SetOssLogConfig(v *DescribeUserConfigsResponseBodyConfigsOssLogConfig) *DescribeUserConfigsResponseBodyConfigs {
	s.OssLogConfig = v
	return s
}

func (s *DescribeUserConfigsResponseBodyConfigs) SetWafConfig(v *DescribeUserConfigsResponseBodyConfigsWafConfig) *DescribeUserConfigsResponseBodyConfigs {
	s.WafConfig = v
	return s
}

func (s *DescribeUserConfigsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeUserConfigsResponseBodyConfigsOssLogConfig struct {
	// The name of the bucket.
	//
	// example:
	//
	// Buckettest
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// Indicates whether the OSS bucket is enabled.
	//
	// example:
	//
	// off
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The prefix.
	//
	// example:
	//
	// test
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s DescribeUserConfigsResponseBodyConfigsOssLogConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConfigsResponseBodyConfigsOssLogConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsResponseBodyConfigsOssLogConfig) GetBucket() *string {
	return s.Bucket
}

func (s *DescribeUserConfigsResponseBodyConfigsOssLogConfig) GetEnable() *string {
	return s.Enable
}

func (s *DescribeUserConfigsResponseBodyConfigsOssLogConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *DescribeUserConfigsResponseBodyConfigsOssLogConfig) SetBucket(v string) *DescribeUserConfigsResponseBodyConfigsOssLogConfig {
	s.Bucket = &v
	return s
}

func (s *DescribeUserConfigsResponseBodyConfigsOssLogConfig) SetEnable(v string) *DescribeUserConfigsResponseBodyConfigsOssLogConfig {
	s.Enable = &v
	return s
}

func (s *DescribeUserConfigsResponseBodyConfigsOssLogConfig) SetPrefix(v string) *DescribeUserConfigsResponseBodyConfigsOssLogConfig {
	s.Prefix = &v
	return s
}

func (s *DescribeUserConfigsResponseBodyConfigsOssLogConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeUserConfigsResponseBodyConfigsWafConfig struct {
	// Indicates whether WAF is enabled.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s DescribeUserConfigsResponseBodyConfigsWafConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConfigsResponseBodyConfigsWafConfig) GoString() string {
	return s.String()
}

func (s *DescribeUserConfigsResponseBodyConfigsWafConfig) GetEnable() *string {
	return s.Enable
}

func (s *DescribeUserConfigsResponseBodyConfigsWafConfig) SetEnable(v string) *DescribeUserConfigsResponseBodyConfigsWafConfig {
	s.Enable = &v
	return s
}

func (s *DescribeUserConfigsResponseBodyConfigsWafConfig) Validate() error {
	return dara.Validate(s)
}
