// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserPropertyValueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetUserPropertyValueResponseBody
	GetRequestId() *string
}

type SetUserPropertyValueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetUserPropertyValueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetUserPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserPropertyValueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetUserPropertyValueResponseBody) SetRequestId(v string) *SetUserPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetUserPropertyValueResponseBody) Validate() error {
	return dara.Validate(s)
}
