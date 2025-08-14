// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSampleDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteSampleDataRequest
	GetLang() *string
	SetId(v string) *DeleteSampleDataRequest
	GetId() *string
	SetRegId(v string) *DeleteSampleDataRequest
	GetRegId() *string
}

type DeleteSampleDataRequest struct {
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
	// Primary key ID
	//
	// example:
	//
	// 3144
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DeleteSampleDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSampleDataRequest) GoString() string {
	return s.String()
}

func (s *DeleteSampleDataRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteSampleDataRequest) GetId() *string {
	return s.Id
}

func (s *DeleteSampleDataRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteSampleDataRequest) SetLang(v string) *DeleteSampleDataRequest {
	s.Lang = &v
	return s
}

func (s *DeleteSampleDataRequest) SetId(v string) *DeleteSampleDataRequest {
	s.Id = &v
	return s
}

func (s *DeleteSampleDataRequest) SetRegId(v string) *DeleteSampleDataRequest {
	s.RegId = &v
	return s
}

func (s *DeleteSampleDataRequest) Validate() error {
	return dara.Validate(s)
}
