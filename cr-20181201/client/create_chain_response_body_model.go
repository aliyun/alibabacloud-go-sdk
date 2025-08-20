// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChainId(v string) *CreateChainResponseBody
	GetChainId() *string
	SetCode(v string) *CreateChainResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CreateChainResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CreateChainResponseBody
	GetRequestId() *string
}

type CreateChainResponseBody struct {
	// The ID of the delivery chain.
	//
	// example:
	//
	// chi-02ymhtwl3cq8****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4BC03B36-E515-5806-99AC-268AE3C0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChainResponseBody) GetChainId() *string {
	return s.ChainId
}

func (s *CreateChainResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateChainResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CreateChainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChainResponseBody) SetChainId(v string) *CreateChainResponseBody {
	s.ChainId = &v
	return s
}

func (s *CreateChainResponseBody) SetCode(v string) *CreateChainResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChainResponseBody) SetIsSuccess(v bool) *CreateChainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateChainResponseBody) SetRequestId(v string) *CreateChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChainResponseBody) Validate() error {
	return dara.Validate(s)
}
