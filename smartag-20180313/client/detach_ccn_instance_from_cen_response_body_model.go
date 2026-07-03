// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachCcnInstanceFromCenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachCcnInstanceFromCenResponseBody
	GetRequestId() *string
}

type DetachCcnInstanceFromCenResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 96AF7326-B6DE-4188-8638-56A6164F62D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachCcnInstanceFromCenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachCcnInstanceFromCenResponseBody) GoString() string {
	return s.String()
}

func (s *DetachCcnInstanceFromCenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachCcnInstanceFromCenResponseBody) SetRequestId(v string) *DetachCcnInstanceFromCenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachCcnInstanceFromCenResponseBody) Validate() error {
	return dara.Validate(s)
}
