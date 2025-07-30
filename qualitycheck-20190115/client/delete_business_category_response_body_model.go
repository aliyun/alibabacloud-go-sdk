// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBusinessCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteBusinessCategoryResponseBody
	GetCode() *string
	SetData(v string) *DeleteBusinessCategoryResponseBody
	GetData() *string
	SetMessage(v string) *DeleteBusinessCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBusinessCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBusinessCategoryResponseBody
	GetSuccess() *bool
}

type DeleteBusinessCategoryResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBusinessCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBusinessCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBusinessCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBusinessCategoryResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteBusinessCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBusinessCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBusinessCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBusinessCategoryResponseBody) SetCode(v string) *DeleteBusinessCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetData(v string) *DeleteBusinessCategoryResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetMessage(v string) *DeleteBusinessCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetRequestId(v string) *DeleteBusinessCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) SetSuccess(v bool) *DeleteBusinessCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBusinessCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
