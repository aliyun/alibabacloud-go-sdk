// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiServerEipResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyApiServerEipResourceResponseBody
	GetRequestId() *string
}

type ModifyApiServerEipResourceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApiServerEipResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiServerEipResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiServerEipResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApiServerEipResourceResponseBody) SetRequestId(v string) *ModifyApiServerEipResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApiServerEipResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
