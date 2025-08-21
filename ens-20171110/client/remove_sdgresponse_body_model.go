// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveSDGResponseBody
	GetRequestId() *string
}

type RemoveSDGResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6666C5A5-75ED-422E-A022-7121FA18C968
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveSDGResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveSDGResponseBody) SetRequestId(v string) *RemoveSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSDGResponseBody) Validate() error {
	return dara.Validate(s)
}
