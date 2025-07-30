// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteDataSetResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteDataSetResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteDataSetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataSetResponseBody
	GetSuccess() *bool
}

type DeleteDataSetResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteDataSetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataSetResponseBody) SetCode(v string) *DeleteDataSetResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetMessage(v string) *DeleteDataSetResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetRequestId(v string) *DeleteDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSetResponseBody) SetSuccess(v bool) *DeleteDataSetResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataSetResponseBody) Validate() error {
	return dara.Validate(s)
}
