// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseEipSegmentAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseEipSegmentAddressResponseBody
	GetRequestId() *string
}

type ReleaseEipSegmentAddressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F7A6301A-64BA-41EC-8284-8F4838C15D1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseEipSegmentAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseEipSegmentAddressResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseEipSegmentAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseEipSegmentAddressResponseBody) SetRequestId(v string) *ReleaseEipSegmentAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseEipSegmentAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
