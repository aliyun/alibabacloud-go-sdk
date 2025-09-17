// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAccessTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableAccessTokenResponseBody
	GetRequestId() *string
}

type DisableAccessTokenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 686BB8A6-BBA5-47E5-8A75-D2ADE433****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAccessTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableAccessTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAccessTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableAccessTokenResponseBody) SetRequestId(v string) *DisableAccessTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableAccessTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
