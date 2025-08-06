// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferIntentionOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *TransferIntentionOwnerResponseBody
	GetData() *bool
	SetErrorCode(v string) *TransferIntentionOwnerResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *TransferIntentionOwnerResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *TransferIntentionOwnerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TransferIntentionOwnerResponseBody
	GetSuccess() *bool
}

type TransferIntentionOwnerResponseBody struct {
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
	// 717711FB-F887-597B-8121-B77437E89B97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TransferIntentionOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferIntentionOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *TransferIntentionOwnerResponseBody) GetData() *bool {
	return s.Data
}

func (s *TransferIntentionOwnerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TransferIntentionOwnerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *TransferIntentionOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferIntentionOwnerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TransferIntentionOwnerResponseBody) SetData(v bool) *TransferIntentionOwnerResponseBody {
	s.Data = &v
	return s
}

func (s *TransferIntentionOwnerResponseBody) SetErrorCode(v string) *TransferIntentionOwnerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TransferIntentionOwnerResponseBody) SetErrorMsg(v string) *TransferIntentionOwnerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *TransferIntentionOwnerResponseBody) SetRequestId(v string) *TransferIntentionOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferIntentionOwnerResponseBody) SetSuccess(v bool) *TransferIntentionOwnerResponseBody {
	s.Success = &v
	return s
}

func (s *TransferIntentionOwnerResponseBody) Validate() error {
	return dara.Validate(s)
}
