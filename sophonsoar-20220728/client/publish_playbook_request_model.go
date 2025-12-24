// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishPlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PublishPlaybookRequest
	GetDescription() *string
	SetPlaybookUuid(v string) *PublishPlaybookRequest
	GetPlaybookUuid() *string
}

type PublishPlaybookRequest struct {
	// The description of the released version.
	//
	// example:
	//
	// This is a waf processing playbook
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The playbook UUID.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~) operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ac343acc-1a61-4084-9a1c-xxxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s PublishPlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishPlaybookRequest) GoString() string {
	return s.String()
}

func (s *PublishPlaybookRequest) GetDescription() *string {
	return s.Description
}

func (s *PublishPlaybookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *PublishPlaybookRequest) SetDescription(v string) *PublishPlaybookRequest {
	s.Description = &v
	return s
}

func (s *PublishPlaybookRequest) SetPlaybookUuid(v string) *PublishPlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *PublishPlaybookRequest) Validate() error {
	return dara.Validate(s)
}
