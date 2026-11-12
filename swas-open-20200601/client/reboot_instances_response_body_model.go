// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootInstancesResponseBody
	GetRequestId() *string
}

type RebootInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RebootInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootInstancesResponseBody) SetRequestId(v string) *RebootInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
