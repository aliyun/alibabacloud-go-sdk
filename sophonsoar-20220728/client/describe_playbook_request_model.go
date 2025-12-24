// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDebugFlag(v int32) *DescribePlaybookRequest
	GetDebugFlag() *int32
	SetLang(v string) *DescribePlaybookRequest
	GetLang() *string
	SetPlaybookUuid(v string) *DescribePlaybookRequest
	GetPlaybookUuid() *string
	SetTaskflowMd5(v string) *DescribePlaybookRequest
	GetTaskflowMd5() *string
}

type DescribePlaybookRequest struct {
	// The flag that indicates whether the playbook is of the debugging or published version. Valid values:
	//
	// 	- **1**: playbook of the debugging version
	//
	// 	- **0**: playbook of the published version
	//
	// example:
	//
	// 0
	DebugFlag *int32 `json:"DebugFlag,omitempty" xml:"DebugFlag,omitempty"`
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
	// 9030076b-6733-4842-b05a-xxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The MD5 hash value of the playbook.
	//
	// example:
	//
	// 7a8f608dc64c242632aa578xxxxx
	TaskflowMd5 *string `json:"TaskflowMd5,omitempty" xml:"TaskflowMd5,omitempty"`
}

func (s DescribePlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlaybookRequest) GoString() string {
	return s.String()
}

func (s *DescribePlaybookRequest) GetDebugFlag() *int32 {
	return s.DebugFlag
}

func (s *DescribePlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePlaybookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribePlaybookRequest) GetTaskflowMd5() *string {
	return s.TaskflowMd5
}

func (s *DescribePlaybookRequest) SetDebugFlag(v int32) *DescribePlaybookRequest {
	s.DebugFlag = &v
	return s
}

func (s *DescribePlaybookRequest) SetLang(v string) *DescribePlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DescribePlaybookRequest) SetPlaybookUuid(v string) *DescribePlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribePlaybookRequest) SetTaskflowMd5(v string) *DescribePlaybookRequest {
	s.TaskflowMd5 = &v
	return s
}

func (s *DescribePlaybookRequest) Validate() error {
	return dara.Validate(s)
}
