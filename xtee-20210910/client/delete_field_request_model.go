// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteFieldRequest
	GetLang() *string
	SetId(v int64) *DeleteFieldRequest
	GetId() *int64
	SetName(v string) *DeleteFieldRequest
	GetName() *string
	SetRegId(v string) *DeleteFieldRequest
	GetRegId() *string
}

type DeleteFieldRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Primary key ID of the field
	//
	// example:
	//
	// 2556
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Name of the field.
	//
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DeleteFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFieldRequest) GoString() string {
	return s.String()
}

func (s *DeleteFieldRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteFieldRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteFieldRequest) GetName() *string {
	return s.Name
}

func (s *DeleteFieldRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteFieldRequest) SetLang(v string) *DeleteFieldRequest {
	s.Lang = &v
	return s
}

func (s *DeleteFieldRequest) SetId(v int64) *DeleteFieldRequest {
	s.Id = &v
	return s
}

func (s *DeleteFieldRequest) SetName(v string) *DeleteFieldRequest {
	s.Name = &v
	return s
}

func (s *DeleteFieldRequest) SetRegId(v string) *DeleteFieldRequest {
	s.RegId = &v
	return s
}

func (s *DeleteFieldRequest) Validate() error {
	return dara.Validate(s)
}
