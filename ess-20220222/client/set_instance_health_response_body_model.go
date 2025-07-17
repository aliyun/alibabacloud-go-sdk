// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceHealthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetInstanceHealthResponseBody
	GetRequestId() *string
}

type SetInstanceHealthResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B755AE57-6093-43E4-938E-DEA422A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetInstanceHealthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceHealthResponseBody) GoString() string {
	return s.String()
}

func (s *SetInstanceHealthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetInstanceHealthResponseBody) SetRequestId(v string) *SetInstanceHealthResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetInstanceHealthResponseBody) Validate() error {
	return dara.Validate(s)
}
