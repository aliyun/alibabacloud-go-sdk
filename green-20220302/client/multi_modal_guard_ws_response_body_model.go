// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardWsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *MultiModalGuardWsResponseBody
	GetRequestId() *string
}

type MultiModalGuardWsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 552F83A7-80C9-17ED-B344-0E13F7D3BF00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MultiModalGuardWsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardWsResponseBody) GoString() string {
	return s.String()
}

func (s *MultiModalGuardWsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MultiModalGuardWsResponseBody) SetRequestId(v string) *MultiModalGuardWsResponseBody {
	s.RequestId = &v
	return s
}

func (s *MultiModalGuardWsResponseBody) Validate() error {
	return dara.Validate(s)
}
