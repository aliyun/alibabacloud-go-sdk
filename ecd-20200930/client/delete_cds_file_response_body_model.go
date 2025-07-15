// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCdsFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCdsFileResponseBody
	GetCode() *string
	SetData(v string) *DeleteCdsFileResponseBody
	GetData() *string
	SetMessage(v string) *DeleteCdsFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCdsFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCdsFileResponseBody
	GetSuccess() *bool
}

type DeleteCdsFileResponseBody struct {
	// The operation result. A value of success indicates that the operation is successful. If the operation failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the data is returned.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5CC5E450-FC43-4F5B-B540-9964BD31****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCdsFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCdsFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCdsFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCdsFileResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteCdsFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCdsFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCdsFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCdsFileResponseBody) SetCode(v string) *DeleteCdsFileResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCdsFileResponseBody) SetData(v string) *DeleteCdsFileResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCdsFileResponseBody) SetMessage(v string) *DeleteCdsFileResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCdsFileResponseBody) SetRequestId(v string) *DeleteCdsFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCdsFileResponseBody) SetSuccess(v bool) *DeleteCdsFileResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCdsFileResponseBody) Validate() error {
	return dara.Validate(s)
}
