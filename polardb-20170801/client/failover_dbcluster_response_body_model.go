// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FailoverDBClusterResponseBody
	GetRequestId() *string
}

type FailoverDBClusterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FailoverDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FailoverDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *FailoverDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FailoverDBClusterResponseBody) SetRequestId(v string) *FailoverDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *FailoverDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
