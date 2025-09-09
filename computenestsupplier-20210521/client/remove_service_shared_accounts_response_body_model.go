// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveServiceSharedAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveServiceSharedAccountsResponseBody
	GetRequestId() *string
}

type RemoveServiceSharedAccountsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveServiceSharedAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveServiceSharedAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveServiceSharedAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveServiceSharedAccountsResponseBody) SetRequestId(v string) *RemoveServiceSharedAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveServiceSharedAccountsResponseBody) Validate() error {
	return dara.Validate(s)
}
