// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOuterAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteOuterAccountResponseBody
	GetCode() *string
	SetData(v bool) *DeleteOuterAccountResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteOuterAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteOuterAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteOuterAccountResponseBody
	GetSuccess() *bool
}

type DeleteOuterAccountResponseBody struct {
	// Status code. A value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation succeeded. Valid values:
	//
	// - **true**: Succeeded
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteOuterAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOuterAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOuterAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteOuterAccountResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteOuterAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteOuterAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOuterAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteOuterAccountResponseBody) SetCode(v string) *DeleteOuterAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOuterAccountResponseBody) SetData(v bool) *DeleteOuterAccountResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteOuterAccountResponseBody) SetMessage(v string) *DeleteOuterAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOuterAccountResponseBody) SetRequestId(v string) *DeleteOuterAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOuterAccountResponseBody) SetSuccess(v bool) *DeleteOuterAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOuterAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
