// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAutoCcWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddAutoCcWhitelistResponseBody
	GetRequestId() *string
}

type AddAutoCcWhitelistResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AB5025DA-5C52-5207-B6AC-3F198758B678
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAutoCcWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAutoCcWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *AddAutoCcWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAutoCcWhitelistResponseBody) SetRequestId(v string) *AddAutoCcWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAutoCcWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
