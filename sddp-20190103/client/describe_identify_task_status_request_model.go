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
	// The ID of the task. Obtain this ID from the Id field in the response from calling the CreateScanTask or ScanOssObjectV1 operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 268
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The language of the request and response. Default value: **zh_cn**. Valid values:
	//
	// - **zh_cn**: Simplified Chinese
	//
	// - **en_us**: U.S. English
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
