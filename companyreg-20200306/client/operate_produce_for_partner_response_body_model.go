// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateProduceForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *OperateProduceForPartnerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *OperateProduceForPartnerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *OperateProduceForPartnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OperateProduceForPartnerResponseBody
	GetSuccess() *bool
}

type OperateProduceForPartnerResponseBody struct {
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// D170A4BA-4528-5E07-B8D5-6449C42EC23F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OperateProduceForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateProduceForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *OperateProduceForPartnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *OperateProduceForPartnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *OperateProduceForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateProduceForPartnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OperateProduceForPartnerResponseBody) SetErrorCode(v string) *OperateProduceForPartnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OperateProduceForPartnerResponseBody) SetErrorMsg(v string) *OperateProduceForPartnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OperateProduceForPartnerResponseBody) SetRequestId(v string) *OperateProduceForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateProduceForPartnerResponseBody) SetSuccess(v bool) *OperateProduceForPartnerResponseBody {
	s.Success = &v
	return s
}

func (s *OperateProduceForPartnerResponseBody) Validate() error {
	return dara.Validate(s)
}
