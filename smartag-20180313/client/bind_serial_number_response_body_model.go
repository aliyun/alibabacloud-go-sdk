// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSerialNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindSerialNumberResponseBody
	GetRequestId() *string
}

type BindSerialNumberResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 61C33444-D8C5-4018-A06C-BA8C8812BEF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindSerialNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindSerialNumberResponseBody) GoString() string {
	return s.String()
}

func (s *BindSerialNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindSerialNumberResponseBody) SetRequestId(v string) *BindSerialNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindSerialNumberResponseBody) Validate() error {
	return dara.Validate(s)
}
