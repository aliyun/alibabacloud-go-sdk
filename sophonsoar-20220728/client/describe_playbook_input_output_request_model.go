// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookInputOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePlaybookInputOutputRequest
	GetLang() *string
	SetPlaybookUuid(v string) *DescribePlaybookInputOutputRequest
	GetPlaybookUuid() *string
}

type DescribePlaybookInputOutputRequest struct {
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
	// b724d2b0-3c3b-4223-9bfd-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookInputOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookInputOutputRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookInputOutputRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePlaybookInputOutputRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribePlaybookInputOutputRequest) SetLang(v string) *DescribePlaybookInputOutputRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookInputOutputRequest) SetPlaybookUuid(v string) *DescribePlaybookInputOutputRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookInputOutputRequest) Validate() error {
	return dara.Validate(s)
}
