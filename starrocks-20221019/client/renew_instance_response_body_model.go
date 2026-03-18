// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RenewInstanceResponseBodyData) *RenewInstanceResponseBody
	GetData() *RenewInstanceResponseBodyData
	SetErrMessage(v string) *RenewInstanceResponseBody
	GetErrMessage() *string
	SetErrorCode(v string) *RenewInstanceResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *RenewInstanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *RenewInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenewInstanceResponseBody
	GetSuccess() *bool
}

type RenewInstanceResponseBody struct {
	Data *RenewInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// Success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE74XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) GetData() *RenewInstanceResponseBodyData {
	return s.Data
}

func (s *RenewInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RenewInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RenewInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RenewInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenewInstanceResponseBody) SetData(v *RenewInstanceResponseBodyData) *RenewInstanceResponseBody {
	s.Data = v
	return s
}

func (s *RenewInstanceResponseBody) SetErrMessage(v string) *RenewInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RenewInstanceResponseBody) SetErrorCode(v string) *RenewInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetHttpStatusCode(v int32) *RenewInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetSuccess(v bool) *RenewInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RenewInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewInstanceResponseBodyData struct {
	OrderIds []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
}

func (s RenewInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RenewInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBodyData) GetOrderIds() []*string {
	return s.OrderIds
}

func (s *RenewInstanceResponseBodyData) SetOrderIds(v []*string) *RenewInstanceResponseBodyData {
	s.OrderIds = v
	return s
}

func (s *RenewInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
