// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSasPmAgentListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSasPmAgentListRequest
	GetLang() *string
	SetUuids(v string) *DescribeSasPmAgentListRequest
	GetUuids() *string
}

type DescribeSasPmAgentListRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The UUID of the server. Separate multiple UUIDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-cb7ae5ee-b2bc-4581-b616-62495f5d****,inet-37cf0e4f-55cc-4b84-8073-b348b4b4****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeSasPmAgentListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSasPmAgentListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSasPmAgentListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSasPmAgentListRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeSasPmAgentListRequest) SetLang(v string) *DescribeSasPmAgentListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSasPmAgentListRequest) SetUuids(v string) *DescribeSasPmAgentListRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeSasPmAgentListRequest) Validate() error {
	return dara.Validate(s)
}
