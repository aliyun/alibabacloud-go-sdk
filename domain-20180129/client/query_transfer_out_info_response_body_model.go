// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTransferOutInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *QueryTransferOutInfoResponseBody
	GetEmail() *string
	SetExpirationDate(v string) *QueryTransferOutInfoResponseBody
	GetExpirationDate() *string
	SetPendingRequestDate(v string) *QueryTransferOutInfoResponseBody
	GetPendingRequestDate() *string
	SetRequestId(v string) *QueryTransferOutInfoResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryTransferOutInfoResponseBody
	GetResultCode() *string
	SetResultMsg(v string) *QueryTransferOutInfoResponseBody
	GetResultMsg() *string
	SetStatus(v int32) *QueryTransferOutInfoResponseBody
	GetStatus() *int32
	SetTransferAuthorizationCodeSendDate(v string) *QueryTransferOutInfoResponseBody
	GetTransferAuthorizationCodeSendDate() *string
}

type QueryTransferOutInfoResponseBody struct {
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 2018-04-13 19:57:56
	ExpirationDate *string `json:"ExpirationDate,omitempty" xml:"ExpirationDate,omitempty"`
	// example:
	//
	// 2018-04-13 19:57:56
	PendingRequestDate *string `json:"PendingRequestDate,omitempty" xml:"PendingRequestDate,omitempty"`
	// example:
	//
	// BBEC5A50-DFDF-482E-8343-B4EB0105E055
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// clientRejected
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// Transfer out rejected
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// 8
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2018-04-13 19:57:56
	TransferAuthorizationCodeSendDate *string `json:"TransferAuthorizationCodeSendDate,omitempty" xml:"TransferAuthorizationCodeSendDate,omitempty"`
}

func (s QueryTransferOutInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTransferOutInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTransferOutInfoResponseBody) GetEmail() *string {
	return s.Email
}

func (s *QueryTransferOutInfoResponseBody) GetExpirationDate() *string {
	return s.ExpirationDate
}

func (s *QueryTransferOutInfoResponseBody) GetPendingRequestDate() *string {
	return s.PendingRequestDate
}

func (s *QueryTransferOutInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTransferOutInfoResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryTransferOutInfoResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryTransferOutInfoResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *QueryTransferOutInfoResponseBody) GetTransferAuthorizationCodeSendDate() *string {
	return s.TransferAuthorizationCodeSendDate
}

func (s *QueryTransferOutInfoResponseBody) SetEmail(v string) *QueryTransferOutInfoResponseBody {
	s.Email = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetExpirationDate(v string) *QueryTransferOutInfoResponseBody {
	s.ExpirationDate = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetPendingRequestDate(v string) *QueryTransferOutInfoResponseBody {
	s.PendingRequestDate = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetRequestId(v string) *QueryTransferOutInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetResultCode(v string) *QueryTransferOutInfoResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetResultMsg(v string) *QueryTransferOutInfoResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetStatus(v int32) *QueryTransferOutInfoResponseBody {
	s.Status = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) SetTransferAuthorizationCodeSendDate(v string) *QueryTransferOutInfoResponseBody {
	s.TransferAuthorizationCodeSendDate = &v
	return s
}

func (s *QueryTransferOutInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
