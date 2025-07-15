// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachPluginResponseBody
	GetRequestId() *string
}

type DetachPluginResponseBody struct {
	// example:
	//
	// AD00F8C0-311B-54A9-ADE2-2436771012DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachPluginResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachPluginResponseBody) SetRequestId(v string) *DetachPluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachPluginResponseBody) Validate() error {
	return dara.Validate(s)
}
