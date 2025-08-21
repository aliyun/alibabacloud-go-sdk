// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbidVsStreamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ForbidVsStreamResponseBody
	GetRequestId() *string
}

type ForbidVsStreamResponseBody struct {
	// example:
	//
	// 119F7639-4646-51A4-B6C1-300D391C0104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ForbidVsStreamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ForbidVsStreamResponseBody) GoString() string {
	return s.String()
}

func (s *ForbidVsStreamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ForbidVsStreamResponseBody) SetRequestId(v string) *ForbidVsStreamResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForbidVsStreamResponseBody) Validate() error {
	return dara.Validate(s)
}
