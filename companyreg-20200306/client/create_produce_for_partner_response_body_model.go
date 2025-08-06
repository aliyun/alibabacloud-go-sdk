// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProduceForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateProduceForPartnerResponseBody
	GetBizId() *string
	SetErrorCode(v string) *CreateProduceForPartnerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateProduceForPartnerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateProduceForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateProduceForPartnerResponseBody
	GetSuccess() *bool
}

type CreateProduceForPartnerResponseBody struct {
	// example:
	//
	// P20210208152920000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 2174AA97-56FB-50FA-B243-0460B9E4CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateProduceForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProduceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProduceForPartnerResponseBody) GetBizId() *string {
	return s.BizId
}

func (s *CreateProduceForPartnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateProduceForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateProduceForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProduceForPartnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProduceForPartnerResponseBody) SetBizId(v string) *CreateProduceForPartnerResponseBody {
	s.BizId = &v
	return s
}

func (s *CreateProduceForPartnerResponseBody) SetErrorCode(v string) *CreateProduceForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProduceForPartnerResponseBody) SetErrorMsg(v string) *CreateProduceForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateProduceForPartnerResponseBody) SetRequestId(v string) *CreateProduceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProduceForPartnerResponseBody) SetSuccess(v bool) *CreateProduceForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProduceForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}
