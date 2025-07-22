// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeRecordTemplatesRequest
	GetAppId() *string
	SetOwnerId(v int64) *DescribeRecordTemplatesRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *DescribeRecordTemplatesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeRecordTemplatesRequest
	GetPageSize() *int32
	SetTemplateIds(v []*string) *DescribeRecordTemplatesRequest
	GetTemplateIds() []*string
}

type DescribeRecordTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 1
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 76dasgb****
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s DescribeRecordTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeRecordTemplatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRecordTemplatesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeRecordTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRecordTemplatesRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *DescribeRecordTemplatesRequest) SetAppId(v string) *DescribeRecordTemplatesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetOwnerId(v int64) *DescribeRecordTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetPageNum(v int32) *DescribeRecordTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetPageSize(v int32) *DescribeRecordTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetTemplateIds(v []*string) *DescribeRecordTemplatesRequest {
	s.TemplateIds = v
	return s
}

func (s *DescribeRecordTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
