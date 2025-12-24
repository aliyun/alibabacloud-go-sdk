// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlaybookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeletePlaybookRequest
	GetLang() *string
	SetPlaybookUuid(v string) *DeletePlaybookRequest
	GetPlaybookUuid() *string
}

type DeletePlaybookRequest struct {
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
	// The UUID of the playbook.
	//
	// >  You can call the [DescribePlaybooks](~~DescribePlaybooks~~)operation to query the playbook UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e99dab31-499b-4307-9248-xxxxxx
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DeletePlaybookRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePlaybookRequest) GoString() string {
	return s.String()
}

func (s *DeletePlaybookRequest) GetLang() *string {
	return s.Lang
}

func (s *DeletePlaybookRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DeletePlaybookRequest) SetLang(v string) *DeletePlaybookRequest {
	s.Lang = &v
	return s
}

func (s *DeletePlaybookRequest) SetPlaybookUuid(v string) *DeletePlaybookRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *DeletePlaybookRequest) Validate() error {
	return dara.Validate(s)
}
