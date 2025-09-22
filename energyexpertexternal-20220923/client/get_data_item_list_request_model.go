// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataItemListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataItemListRequest
	GetCode() *string
}

type GetDataItemListRequest struct {
	// The enterprise code.
	//
	// This parameter is required.
	//
	// example:
	//
	// C-202302-01
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
}

func (s GetDataItemListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataItemListRequest) GoString() string {
	return s.String()
}

func (s *GetDataItemListRequest) GetCode() *string {
	return s.Code
}

func (s *GetDataItemListRequest) SetCode(v string) *GetDataItemListRequest {
	s.Code = &v
	return s
}

func (s *GetDataItemListRequest) Validate() error {
	return dara.Validate(s)
}
