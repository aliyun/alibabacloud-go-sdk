// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBizUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateBizUnitResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateBizUnitResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateBizUnitResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBizUnitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBizUnitResponseBody
	GetSuccess() *bool
}

type UpdateBizUnitResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBizUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBizUnitResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBizUnitResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateBizUnitResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBizUnitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBizUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBizUnitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBizUnitResponseBody) SetCode(v string) *UpdateBizUnitResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBizUnitResponseBody) SetHttpStatusCode(v int32) *UpdateBizUnitResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBizUnitResponseBody) SetMessage(v string) *UpdateBizUnitResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBizUnitResponseBody) SetRequestId(v string) *UpdateBizUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBizUnitResponseBody) SetSuccess(v bool) *UpdateBizUnitResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBizUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
