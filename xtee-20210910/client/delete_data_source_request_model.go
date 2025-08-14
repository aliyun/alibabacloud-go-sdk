// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteDataSourceRequest
	GetLang() *string
	SetId(v int64) *DeleteDataSourceRequest
	GetId() *int64
	SetRegId(v string) *DeleteDataSourceRequest
	GetRegId() *string
}

type DeleteDataSourceRequest struct {
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
	// Primary key ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
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

func (s DeleteDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDataSourceRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDataSourceRequest) GetRegId() *string {
	return s.RegId
}

func (s *DeleteDataSourceRequest) SetLang(v string) *DeleteDataSourceRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDataSourceRequest) SetId(v int64) *DeleteDataSourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteDataSourceRequest) SetRegId(v string) *DeleteDataSourceRequest {
	s.RegId = &v
	return s
}

func (s *DeleteDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
