// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSnatIpForSnatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopSnatIpForSnatEntryResponseBody
	GetRequestId() *string
}

type StopSnatIpForSnatEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9A415CB9-7591-566F-924B-32709578756B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopSnatIpForSnatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopSnatIpForSnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *StopSnatIpForSnatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopSnatIpForSnatEntryResponseBody) SetRequestId(v string) *StopSnatIpForSnatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSnatIpForSnatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
