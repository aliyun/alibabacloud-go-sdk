// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLatestRecordSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeLatestRecordSchemaRequest
	GetLang() *string
	SetPlaybookUuid(v string) *DescribeLatestRecordSchemaRequest
	GetPlaybookUuid() *string
}

type DescribeLatestRecordSchemaRequest struct {
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
	// c5c88b5e-97ca-435d-8c20-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeLatestRecordSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLatestRecordSchemaRequest) GoString() string {
	return s.String()
}

func (s *DescribeLatestRecordSchemaRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeLatestRecordSchemaRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribeLatestRecordSchemaRequest) SetLang(v string) *DescribeLatestRecordSchemaRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLatestRecordSchemaRequest) SetPlaybookUuid(v string) *DescribeLatestRecordSchemaRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribeLatestRecordSchemaRequest) Validate() error {
	return dara.Validate(s)
}
