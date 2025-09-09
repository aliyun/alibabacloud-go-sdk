// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdentifyTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DescribeIdentifyTaskStatusRequest
	GetId() *int64
	SetLang(v string) *DescribeIdentifyTaskStatusRequest
	GetLang() *string
}

type DescribeIdentifyTaskStatusRequest struct {
	// Task ID, obtained from the ID field in the response after calling CreateScanTask or ScanOssObjectV1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 268
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Language type for request and response messages, default is **zh_cn**. Values:
	//
	// - **zh_cn**: Chinese (Simplified)
	//
	// - **en_us**: English (United States)
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeIdentifyTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdentifyTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeIdentifyTaskStatusRequest) GetId() *int64 {
	return s.Id
}

func (s *DescribeIdentifyTaskStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeIdentifyTaskStatusRequest) SetId(v int64) *DescribeIdentifyTaskStatusRequest {
	s.Id = &v
	return s
}

func (s *DescribeIdentifyTaskStatusRequest) SetLang(v string) *DescribeIdentifyTaskStatusRequest {
	s.Lang = &v
	return s
}

func (s *DescribeIdentifyTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
