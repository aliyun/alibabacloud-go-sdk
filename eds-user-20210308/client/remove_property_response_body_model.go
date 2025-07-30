// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemovePropertyResponseBody
	GetRequestId() *string
}

type RemovePropertyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemovePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemovePropertyResponseBody) SetRequestId(v string) *RemovePropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePropertyResponseBody) Validate() error {
	return dara.Validate(s)
}
