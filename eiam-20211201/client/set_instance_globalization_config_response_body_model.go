// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceGlobalizationConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetInstanceGlobalizationConfigResponseBody
	GetRequestId() *string
}

type SetInstanceGlobalizationConfigResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstanceGlobalizationConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceGlobalizationConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstanceGlobalizationConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetInstanceGlobalizationConfigResponseBody) SetRequestId(v string) *SetInstanceGlobalizationConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetInstanceGlobalizationConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
