// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseTrafficControlTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseTrafficControlTaskResponseBody
	GetRequestId() *string
}

type ReleaseTrafficControlTaskResponseBody struct {
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseTrafficControlTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseTrafficControlTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseTrafficControlTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseTrafficControlTaskResponseBody) SetRequestId(v string) *ReleaseTrafficControlTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseTrafficControlTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
