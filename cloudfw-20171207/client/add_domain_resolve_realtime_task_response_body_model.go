// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainResolveRealtimeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddDomainResolveRealtimeTaskResponseBody
	GetRequestId() *string
}

type AddDomainResolveRealtimeTaskResponseBody struct {
	// example:
	//
	// 337A4DBA-8A01-5E9C-99CA-84293E13****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDomainResolveRealtimeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDomainResolveRealtimeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AddDomainResolveRealtimeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDomainResolveRealtimeTaskResponseBody) SetRequestId(v string) *AddDomainResolveRealtimeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDomainResolveRealtimeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
