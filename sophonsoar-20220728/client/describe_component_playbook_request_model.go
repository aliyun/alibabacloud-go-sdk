// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeComponentPlaybookRequest
	GetLang() *string
	SetPlaybookUuid(v string) *DescribeComponentPlaybookRequest
	GetPlaybookUuid() *string
}

type DescribeComponentPlaybookRequest struct {
	// The language of the content within the request and the response. Valid values:
	//
	// 	- **zh*	- (default): Chinese
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
	// ac343acc-1a61-4084-9a1cxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeComponentPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentPlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentPlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeComponentPlaybookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribeComponentPlaybookRequest) SetLang(v string) *DescribeComponentPlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DescribeComponentPlaybookRequest) SetPlaybookUuid(v string) *DescribeComponentPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribeComponentPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
