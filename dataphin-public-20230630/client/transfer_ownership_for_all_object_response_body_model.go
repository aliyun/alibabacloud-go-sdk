// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferOwnershipForAllObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TransferOwnershipForAllObjectResponseBody
	GetCode() *string
	SetData(v int64) *TransferOwnershipForAllObjectResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *TransferOwnershipForAllObjectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *TransferOwnershipForAllObjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *TransferOwnershipForAllObjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TransferOwnershipForAllObjectResponseBody
	GetSuccess() *bool
}

type TransferOwnershipForAllObjectResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s TransferOwnershipForAllObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferOwnershipForAllObjectResponseBody) GoString() string {
	return s.String()
}

func (s *TransferOwnershipForAllObjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *TransferOwnershipForAllObjectResponseBody) GetData() *int64 {
	return s.Data
}

func (s *TransferOwnershipForAllObjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *TransferOwnershipForAllObjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TransferOwnershipForAllObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferOwnershipForAllObjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TransferOwnershipForAllObjectResponseBody) SetCode(v string) *TransferOwnershipForAllObjectResponseBody {
	s.Code = &v
	return s
}

func (s *TransferOwnershipForAllObjectResponseBody) SetData(v int64) *TransferOwnershipForAllObjectResponseBody {
	s.Data = &v
	return s
}

func (s *TransferOwnershipForAllObjectResponseBody) SetHttpStatusCode(v int32) *TransferOwnershipForAllObjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TransferOwnershipForAllObjectResponseBody) SetMessage(v string) *TransferOwnershipForAllObjectResponseBody {
	s.Message = &v
	return s
}

func (s *TransferOwnershipForAllObjectResponseBody) SetRequestId(v string) *TransferOwnershipForAllObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferOwnershipForAllObjectResponseBody) SetSuccess(v bool) *TransferOwnershipForAllObjectResponseBody {
	s.Success = &v
	return s
}

func (s *TransferOwnershipForAllObjectResponseBody) Validate() error {
	return dara.Validate(s)
}
