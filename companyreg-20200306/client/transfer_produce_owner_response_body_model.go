// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferProduceOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *TransferProduceOwnerResponseBody
	GetData() *bool
	SetErrorCode(v string) *TransferProduceOwnerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TransferProduceOwnerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TransferProduceOwnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TransferProduceOwnerResponseBody
	GetSuccess() *bool
}

type TransferProduceOwnerResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// DD5639FF-1240-51DE-9BA8-2075670A1EAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferProduceOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferProduceOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *TransferProduceOwnerResponseBody) GetData() *bool {
	return s.Data
}

func (s *TransferProduceOwnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TransferProduceOwnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TransferProduceOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferProduceOwnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TransferProduceOwnerResponseBody) SetData(v bool) *TransferProduceOwnerResponseBody {
	s.Data = &v
	return s
}

func (s *TransferProduceOwnerResponseBody) SetErrorCode(v string) *TransferProduceOwnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TransferProduceOwnerResponseBody) SetErrorMsg(v string) *TransferProduceOwnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TransferProduceOwnerResponseBody) SetRequestId(v string) *TransferProduceOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferProduceOwnerResponseBody) SetSuccess(v bool) *TransferProduceOwnerResponseBody {
	s.Success = &v
	return s
}

func (s *TransferProduceOwnerResponseBody) Validate() error {
	return dara.Validate(s)
}
