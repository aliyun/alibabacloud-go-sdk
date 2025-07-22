// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFundAccountTransferResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetadata(v interface{}) *CreateFundAccountTransferResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *CreateFundAccountTransferResponseBody
	GetRequestId() *string
}

type CreateFundAccountTransferResponseBody struct {
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 1BB79-5B23-3EA-BB4F-352F93E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFundAccountTransferResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFundAccountTransferResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFundAccountTransferResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *CreateFundAccountTransferResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFundAccountTransferResponseBody) SetMetadata(v interface{}) *CreateFundAccountTransferResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateFundAccountTransferResponseBody) SetRequestId(v string) *CreateFundAccountTransferResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFundAccountTransferResponseBody) Validate() error {
	return dara.Validate(s)
}
