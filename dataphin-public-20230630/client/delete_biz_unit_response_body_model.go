// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBizUnitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteBizUnitResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteBizUnitResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteBizUnitResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteBizUnitResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBizUnitResponseBody
	GetSuccess() *bool
}

type DeleteBizUnitResponseBody struct {
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

func (s DeleteBizUnitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBizUnitResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBizUnitResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteBizUnitResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteBizUnitResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteBizUnitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBizUnitResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBizUnitResponseBody) SetCode(v string) *DeleteBizUnitResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteBizUnitResponseBody) SetHttpStatusCode(v int32) *DeleteBizUnitResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBizUnitResponseBody) SetMessage(v string) *DeleteBizUnitResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBizUnitResponseBody) SetRequestId(v string) *DeleteBizUnitResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBizUnitResponseBody) SetSuccess(v bool) *DeleteBizUnitResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBizUnitResponseBody) Validate() error {
	return dara.Validate(s)
}
