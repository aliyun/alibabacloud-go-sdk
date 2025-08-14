// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNameListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteNameListRequest
	GetLang() *string
	SetIds(v string) *DeleteNameListRequest
	GetIds() *string
	SetRegId(v string) *DeleteNameListRequest
	GetRegId() *string
}

type DeleteNameListRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Name list ID, multiple IDs should be separated by commas
	//
	// example:
	//
	// 23,24,25
	Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DeleteNameListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNameListRequest) GoString() string {
	return s.String()
}

func (s *DeleteNameListRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteNameListRequest) GetIds() *string {
	return s.Ids
}

func (s *DeleteNameListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteNameListRequest) SetLang(v string) *DeleteNameListRequest {
	s.Lang = &v
	return s
}

func (s *DeleteNameListRequest) SetIds(v string) *DeleteNameListRequest {
	s.Ids = &v
	return s
}

func (s *DeleteNameListRequest) SetRegId(v string) *DeleteNameListRequest {
	s.RegId = &v
	return s
}

func (s *DeleteNameListRequest) Validate() error {
	return dara.Validate(s)
}
