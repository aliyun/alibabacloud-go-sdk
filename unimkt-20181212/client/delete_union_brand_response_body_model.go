// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUnionBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteUnionBrandResponseBody
	GetCode() *string
	SetData(v bool) *DeleteUnionBrandResponseBody
	GetData() *bool
	SetErrorMsg(v string) *DeleteUnionBrandResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *DeleteUnionBrandResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteUnionBrandResponseBody
	GetSuccess() *string
}

type DeleteUnionBrandResponseBody struct {
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// error check permissions
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6CBD8D0F-D84A-5545-817C-16C7430FD612
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteUnionBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUnionBrandResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUnionBrandResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteUnionBrandResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteUnionBrandResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DeleteUnionBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteUnionBrandResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteUnionBrandResponseBody) SetCode(v string) *DeleteUnionBrandResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUnionBrandResponseBody) SetData(v bool) *DeleteUnionBrandResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteUnionBrandResponseBody) SetErrorMsg(v string) *DeleteUnionBrandResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteUnionBrandResponseBody) SetRequestId(v string) *DeleteUnionBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteUnionBrandResponseBody) SetSuccess(v string) *DeleteUnionBrandResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUnionBrandResponseBody) Validate() error {
	return dara.Validate(s)
}
