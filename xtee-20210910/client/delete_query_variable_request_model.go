// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueryVariableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteQueryVariableRequest
	GetLang() *string
	SetId(v int64) *DeleteQueryVariableRequest
	GetId() *int64
	SetRegId(v string) *DeleteQueryVariableRequest
	GetRegId() *string
}

type DeleteQueryVariableRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Variable ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DeleteQueryVariableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueryVariableRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueryVariableRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteQueryVariableRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteQueryVariableRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteQueryVariableRequest) SetLang(v string) *DeleteQueryVariableRequest {
	s.Lang = &v
	return s
}

func (s *DeleteQueryVariableRequest) SetId(v int64) *DeleteQueryVariableRequest {
	s.Id = &v
	return s
}

func (s *DeleteQueryVariableRequest) SetRegId(v string) *DeleteQueryVariableRequest {
	s.RegId = &v
	return s
}

func (s *DeleteQueryVariableRequest) Validate() error {
	return dara.Validate(s)
}
