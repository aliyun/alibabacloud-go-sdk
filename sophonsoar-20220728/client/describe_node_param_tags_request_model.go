// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNodeParamTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNodeParamTagsRequest
	GetLang() *string
	SetNodeName(v string) *DescribeNodeParamTagsRequest
	GetNodeName() *string
	SetPlaybookUuid(v string) *DescribeNodeParamTagsRequest
	GetPlaybookUuid() *string
}

type DescribeNodeParamTagsRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// python3_2
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac343acc-1a61-4084-9a1c-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeNodeParamTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeParamTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeParamTagsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNodeParamTagsRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeNodeParamTagsRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribeNodeParamTagsRequest) SetLang(v string) *DescribeNodeParamTagsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNodeParamTagsRequest) SetNodeName(v string) *DescribeNodeParamTagsRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeNodeParamTagsRequest) SetPlaybookUuid(v string) *DescribeNodeParamTagsRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribeNodeParamTagsRequest) Validate() error {
	return dara.Validate(s)
}
