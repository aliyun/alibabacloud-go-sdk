// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdsFileShareLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCdsFileShareLinkResponseBody
	GetCode() *string
	SetData(v *CdsFileShareLinkModel) *CreateCdsFileShareLinkResponseBody
	GetData() *CdsFileShareLinkModel
	SetMessage(v string) *CreateCdsFileShareLinkResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCdsFileShareLinkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCdsFileShareLinkResponseBody
	GetSuccess() *bool
}

type CreateCdsFileShareLinkResponseBody struct {
	// The operation result. The value success indicates that the operation is successful. If the operation failed, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data information.
	Data *CdsFileShareLinkModel `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message that is returned. This parameter is not returned if the value of Code is success.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 51592A88-0F2C-55E6-AD2C-2AD9C10D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
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

func (s CreateCdsFileShareLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCdsFileShareLinkResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCdsFileShareLinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCdsFileShareLinkResponseBody) GetData() *CdsFileShareLinkModel {
	return s.Data
}

func (s *CreateCdsFileShareLinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCdsFileShareLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCdsFileShareLinkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCdsFileShareLinkResponseBody) SetCode(v string) *CreateCdsFileShareLinkResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCdsFileShareLinkResponseBody) SetData(v *CdsFileShareLinkModel) *CreateCdsFileShareLinkResponseBody {
	s.Data = v
	return s
}

func (s *CreateCdsFileShareLinkResponseBody) SetMessage(v string) *CreateCdsFileShareLinkResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCdsFileShareLinkResponseBody) SetRequestId(v string) *CreateCdsFileShareLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCdsFileShareLinkResponseBody) SetSuccess(v bool) *CreateCdsFileShareLinkResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCdsFileShareLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
