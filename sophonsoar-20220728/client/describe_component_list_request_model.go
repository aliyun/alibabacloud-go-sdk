// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeComponentListRequest
	GetLang() *string
	SetPlaybookUuid(v string) *DescribeComponentListRequest
	GetPlaybookUuid() *string
}

type DescribeComponentListRequest struct {
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
	// b724d2b0-3c3b-4223-9bfd-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeComponentListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentListRequest) GoString() string {
	return s.String()
}

func (s *DescribeComponentListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeComponentListRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribeComponentListRequest) SetLang(v string) *DescribeComponentListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeComponentListRequest) SetPlaybookUuid(v string) *DescribeComponentListRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribeComponentListRequest) Validate() error {
	return dara.Validate(s)
}
