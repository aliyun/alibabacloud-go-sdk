// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistinctReleasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDistinctReleasesRequest
	GetLang() *string
	SetPlaybookUuid(v string) *DescribeDistinctReleasesRequest
	GetPlaybookUuid() *string
	SetTaskflowMd5(v string) *DescribeDistinctReleasesRequest
	GetTaskflowMd5() *string
}

type DescribeDistinctReleasesRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bc0b8424-535c-4ed5-bd94-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The MD5 value of the playbook XML configuration.
	//
	// example:
	//
	// be0a4ef084dd174abe47xxxxx
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
}

func (s DescribeDistinctReleasesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistinctReleasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistinctReleasesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDistinctReleasesRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribeDistinctReleasesRequest) GetTaskflowMd5() *string {
	return s.TaskflowMd5
}

func (s *DescribeDistinctReleasesRequest) SetLang(v string) *DescribeDistinctReleasesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDistinctReleasesRequest) SetPlaybookUuid(v string) *DescribeDistinctReleasesRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribeDistinctReleasesRequest) SetTaskflowMd5(v string) *DescribeDistinctReleasesRequest {
	s.TaskflowMd5 = &v
	return s
}

func (s *DescribeDistinctReleasesRequest) Validate() error {
	return dara.Validate(s)
}
