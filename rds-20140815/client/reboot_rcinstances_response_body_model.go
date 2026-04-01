// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRCInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootRCInstancesResponseBody
	GetRequestId() *string
}

type RebootRCInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 32A5B40E-50DA-5166-9B22-35F00C5D1BC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootRCInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootRCInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RebootRCInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootRCInstancesResponseBody) SetRequestId(v string) *RebootRCInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootRCInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
