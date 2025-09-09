// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssObjectDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DescribeOssObjectDetailRequest
	GetId() *int64
	SetLang(v string) *DescribeOssObjectDetailRequest
	GetLang() *string
}

type DescribeOssObjectDetailRequest struct {
	// The ID of the OSS object.
	//
	// >  You can call the [DescribeOssObjects](https://help.aliyun.com/document_detail/410152.html) operation to obtain the ID of the OSS object.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345213
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeOssObjectDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeOssObjectDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOssObjectDetailRequest) SetId(v int64) *DescribeOssObjectDetailRequest {
	s.Id = &v
	return s
}

func (s *DescribeOssObjectDetailRequest) SetLang(v string) *DescribeOssObjectDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOssObjectDetailRequest) Validate() error {
	return dara.Validate(s)
}
