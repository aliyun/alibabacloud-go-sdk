// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCdsFileShareLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelCdsFileShareLinkResponseBody
	GetCode() *string
	SetData(v bool) *CancelCdsFileShareLinkResponseBody
	GetData() *bool
	SetMessage(v string) *CancelCdsFileShareLinkResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelCdsFileShareLinkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelCdsFileShareLinkResponseBody
	GetSuccess() *bool
}

type CancelCdsFileShareLinkResponseBody struct {
	// The operation result. The value success indicates that the operation is successful. If the operation failed, an error message is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data information.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message that is returned if the request failed. This parameter is not returned if the value of Code is `success`.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
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

func (s CancelCdsFileShareLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCdsFileShareLinkResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCdsFileShareLinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelCdsFileShareLinkResponseBody) GetData() *bool {
	return s.Data
}

func (s *CancelCdsFileShareLinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelCdsFileShareLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCdsFileShareLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelCdsFileShareLinkResponseBody) SetCode(v string) *CancelCdsFileShareLinkResponseBody {
	s.Code = &v
	return s
}

func (s *CancelCdsFileShareLinkResponseBody) SetData(v bool) *CancelCdsFileShareLinkResponseBody {
	s.Data = &v
	return s
}

func (s *CancelCdsFileShareLinkResponseBody) SetMessage(v string) *CancelCdsFileShareLinkResponseBody {
	s.Message = &v
	return s
}

func (s *CancelCdsFileShareLinkResponseBody) SetRequestId(v string) *CancelCdsFileShareLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCdsFileShareLinkResponseBody) SetSuccess(v bool) *CancelCdsFileShareLinkResponseBody {
	s.Success = &v
	return s
}

func (s *CancelCdsFileShareLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
