// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePayOrderToMsenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMpaasOrderCreateResponse(v *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse) *CreatePayOrderToMsenceResponseBody
	GetMpaasOrderCreateResponse() *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse
	SetRequestId(v string) *CreatePayOrderToMsenceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *CreatePayOrderToMsenceResponseBody
	GetResultCode() *string
	SetResultMsg(v string) *CreatePayOrderToMsenceResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *CreatePayOrderToMsenceResponseBody
	GetSuccess() *bool
}

type CreatePayOrderToMsenceResponseBody struct {
	MpaasOrderCreateResponse *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse `json:"MpaasOrderCreateResponse,omitempty" xml:"MpaasOrderCreateResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// SUCCESS
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePayOrderToMsenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePayOrderToMsenceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePayOrderToMsenceResponseBody) GetMpaasOrderCreateResponse() *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse {
	return s.MpaasOrderCreateResponse
}

func (s *CreatePayOrderToMsenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePayOrderToMsenceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *CreatePayOrderToMsenceResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *CreatePayOrderToMsenceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePayOrderToMsenceResponseBody) SetMpaasOrderCreateResponse(v *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse) *CreatePayOrderToMsenceResponseBody {
	s.MpaasOrderCreateResponse = v
	return s
}

func (s *CreatePayOrderToMsenceResponseBody) SetRequestId(v string) *CreatePayOrderToMsenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePayOrderToMsenceResponseBody) SetResultCode(v string) *CreatePayOrderToMsenceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *CreatePayOrderToMsenceResponseBody) SetResultMsg(v string) *CreatePayOrderToMsenceResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *CreatePayOrderToMsenceResponseBody) SetSuccess(v bool) *CreatePayOrderToMsenceResponseBody {
	s.Success = &v
	return s
}

func (s *CreatePayOrderToMsenceResponseBody) Validate() error {
	if s.MpaasOrderCreateResponse != nil {
		if err := s.MpaasOrderCreateResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse struct {
	// example:
	//
	// 3929520
	BizOrderId *string `json:"BizOrderId,omitempty" xml:"BizOrderId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse) GoString() string {
	return s.String()
}

func (s *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse) GetBizOrderId() *string {
	return s.BizOrderId
}

func (s *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse) GetSuccess() *bool {
	return s.Success
}

func (s *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse) SetBizOrderId(v string) *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse {
	s.BizOrderId = &v
	return s
}

func (s *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse) SetSuccess(v bool) *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse {
	s.Success = &v
	return s
}

func (s *CreatePayOrderToMsenceResponseBodyMpaasOrderCreateResponse) Validate() error {
	return dara.Validate(s)
}
