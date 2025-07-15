// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLiveDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopLiveDomainResponseBody
	GetRequestId() *string
}

type StopLiveDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 324AEFFF-308C-4DA7-8CD3-01B277B98F28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopLiveDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopLiveDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StopLiveDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopLiveDomainResponseBody) SetRequestId(v string) *StopLiveDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopLiveDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
