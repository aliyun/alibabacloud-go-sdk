// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCopyImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelCopyImageResponseBody
	GetRequestId() *string
}

type CancelCopyImageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5BEFE642-A383-4A18-8939-FB7DE452****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelCopyImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCopyImageResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCopyImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCopyImageResponseBody) SetRequestId(v string) *CancelCopyImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCopyImageResponseBody) Validate() error {
	return dara.Validate(s)
}
