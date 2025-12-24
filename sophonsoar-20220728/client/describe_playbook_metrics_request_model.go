// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePlaybookMetricsRequest
	GetLang() *string
	SetPlaybookUuid(v string) *DescribePlaybookMetricsRequest
	GetPlaybookUuid() *string
}

type DescribePlaybookMetricsRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2a687089-d4dd-47d4-9709-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookMetricsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePlaybookMetricsRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribePlaybookMetricsRequest) SetLang(v string) *DescribePlaybookMetricsRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookMetricsRequest) SetPlaybookUuid(v string) *DescribePlaybookMetricsRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookMetricsRequest) Validate() error {
	return dara.Validate(s)
}
