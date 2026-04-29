// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleObjectScanEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *HandleObjectScanEventResponseBody
	GetRequestId() *string
}

type HandleObjectScanEventResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HandleObjectScanEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HandleObjectScanEventResponseBody) GoString() string {
	return s.String()
}

func (s *HandleObjectScanEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HandleObjectScanEventResponseBody) SetRequestId(v string) *HandleObjectScanEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleObjectScanEventResponseBody) Validate() error {
	return dara.Validate(s)
}
