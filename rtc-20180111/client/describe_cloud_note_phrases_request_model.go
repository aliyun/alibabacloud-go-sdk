// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudNotePhrasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeCloudNotePhrasesRequest
	GetAppId() *string
	SetCondition(v *DescribeCloudNotePhrasesRequestCondition) *DescribeCloudNotePhrasesRequest
	GetCondition() *DescribeCloudNotePhrasesRequestCondition
	SetPageNum(v int32) *DescribeCloudNotePhrasesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeCloudNotePhrasesRequest
	GetPageSize() *int32
}

type DescribeCloudNotePhrasesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId     *string                                   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Condition *DescribeCloudNotePhrasesRequestCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeCloudNotePhrasesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotePhrasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotePhrasesRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeCloudNotePhrasesRequest) GetCondition() *DescribeCloudNotePhrasesRequestCondition {
	return s.Condition
}

func (s *DescribeCloudNotePhrasesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeCloudNotePhrasesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudNotePhrasesRequest) SetAppId(v string) *DescribeCloudNotePhrasesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCloudNotePhrasesRequest) SetCondition(v *DescribeCloudNotePhrasesRequestCondition) *DescribeCloudNotePhrasesRequest {
	s.Condition = v
	return s
}

func (s *DescribeCloudNotePhrasesRequest) SetPageNum(v int32) *DescribeCloudNotePhrasesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeCloudNotePhrasesRequest) SetPageSize(v int32) *DescribeCloudNotePhrasesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudNotePhrasesRequest) Validate() error {
	if s.Condition != nil {
		if err := s.Condition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudNotePhrasesRequestCondition struct {
	// example:
	//
	// ac7N****112121
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCloudNotePhrasesRequestCondition) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudNotePhrasesRequestCondition) GoString() string {
	return s.String()
}

func (s *DescribeCloudNotePhrasesRequestCondition) GetId() *string {
	return s.Id
}

func (s *DescribeCloudNotePhrasesRequestCondition) GetName() *string {
	return s.Name
}

func (s *DescribeCloudNotePhrasesRequestCondition) SetId(v string) *DescribeCloudNotePhrasesRequestCondition {
	s.Id = &v
	return s
}

func (s *DescribeCloudNotePhrasesRequestCondition) SetName(v string) *DescribeCloudNotePhrasesRequestCondition {
	s.Name = &v
	return s
}

func (s *DescribeCloudNotePhrasesRequestCondition) Validate() error {
	return dara.Validate(s)
}
