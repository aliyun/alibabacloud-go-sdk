// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveImageResponseBody
	GetRequestId() *string
}

type RemoveImageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveImageResponseBody) SetRequestId(v string) *RemoveImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveImageResponseBody) Validate() error {
	return dara.Validate(s)
}
