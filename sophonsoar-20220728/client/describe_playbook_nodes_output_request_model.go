// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookNodesOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePlaybookNodesOutputRequest
	GetLang() *string
	SetNodeName(v string) *DescribePlaybookNodesOutputRequest
	GetNodeName() *string
	SetPlaybookUuid(v string) *DescribePlaybookNodesOutputRequest
	GetPlaybookUuid() *string
}

type DescribePlaybookNodesOutputRequest struct {
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
	// The name of the component node.
	//
	// This parameter is required.
	//
	// example:
	//
	// DataFormat_1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the UUIDs of playbooks.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac343acc-1a61-4084-9a1c-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribePlaybookNodesOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookNodesOutputRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookNodesOutputRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePlaybookNodesOutputRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribePlaybookNodesOutputRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribePlaybookNodesOutputRequest) SetLang(v string) *DescribePlaybookNodesOutputRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookNodesOutputRequest) SetNodeName(v string) *DescribePlaybookNodesOutputRequest {
	s.NodeName = &v
	return s
}

func (s *DescribePlaybookNodesOutputRequest) SetPlaybookUuid(v string) *DescribePlaybookNodesOutputRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookNodesOutputRequest) Validate() error {
	return dara.Validate(s)
}
