// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVerifyResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteVerifyResultResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteVerifyResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVerifyResultResponseBody
	GetRequestId() *string
	SetResult(v *DeleteVerifyResultResponseBodyResult) *DeleteVerifyResultResponseBody
	GetResult() *DeleteVerifyResultResponseBodyResult
}

type DeleteVerifyResultResponseBody struct {
	// The return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EB35****87EBA1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result *DeleteVerifyResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DeleteVerifyResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVerifyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteVerifyResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVerifyResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVerifyResultResponseBody) GetResult() *DeleteVerifyResultResponseBodyResult {
	return s.Result
}

func (s *DeleteVerifyResultResponseBody) SetCode(v string) *DeleteVerifyResultResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVerifyResultResponseBody) SetMessage(v string) *DeleteVerifyResultResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVerifyResultResponseBody) SetRequestId(v string) *DeleteVerifyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVerifyResultResponseBody) SetResult(v *DeleteVerifyResultResponseBodyResult) *DeleteVerifyResultResponseBody {
	s.Result = v
	return s
}

func (s *DeleteVerifyResultResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteVerifyResultResponseBodyResult struct {
	// The deletion result. Valid values:
	//
	// - Y: The deletion was successful.
	//
	// - N: The deletion failed.
	//
	// example:
	//
	// Y/N
	DeleteResult *string `json:"DeleteResult,omitempty" xml:"DeleteResult,omitempty"`
	// The unique identifier of the authentication request.
	//
	// example:
	//
	// 4ab0b***cbde97
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s DeleteVerifyResultResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DeleteVerifyResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DeleteVerifyResultResponseBodyResult) GetDeleteResult() *string {
	return s.DeleteResult
}

func (s *DeleteVerifyResultResponseBodyResult) GetTransactionId() *string {
	return s.TransactionId
}

func (s *DeleteVerifyResultResponseBodyResult) SetDeleteResult(v string) *DeleteVerifyResultResponseBodyResult {
	s.DeleteResult = &v
	return s
}

func (s *DeleteVerifyResultResponseBodyResult) SetTransactionId(v string) *DeleteVerifyResultResponseBodyResult {
	s.TransactionId = &v
	return s
}

func (s *DeleteVerifyResultResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
