// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSelfBindVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteSelfBindVariableRequest
	GetLang() *string
	SetId(v int64) *DeleteSelfBindVariableRequest
	GetId() *int64
	SetRegId(v string) *DeleteSelfBindVariableRequest
	GetRegId() *string
}

type DeleteSelfBindVariableRequest struct {
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
	// Variable ID
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DeleteSelfBindVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSelfBindVariableRequest) GoString() string {
	return s.String()
}

func (s *DeleteSelfBindVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteSelfBindVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteSelfBindVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteSelfBindVariableRequest) SetLang(v string) *DeleteSelfBindVariableRequest {
	s.Lang = &v
	return s
}

func (s *DeleteSelfBindVariableRequest) SetId(v int64) *DeleteSelfBindVariableRequest {
	s.Id = &v
	return s
}

func (s *DeleteSelfBindVariableRequest) SetRegId(v string) *DeleteSelfBindVariableRequest {
	s.RegId = &v
	return s
}

func (s *DeleteSelfBindVariableRequest) Validate() error {
	return dara.Validate(s)
}
