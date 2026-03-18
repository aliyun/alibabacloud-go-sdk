// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachToPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachToPolicyResponseBody
	GetRequestId() *string
}

type AttachToPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DC245DEE-9800-5579-BF99-189D6A5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachToPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachToPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AttachToPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachToPolicyResponseBody) SetRequestId(v string) *AttachToPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachToPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
