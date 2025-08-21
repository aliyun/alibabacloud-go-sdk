// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHeadersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHeadersResponseBody
	GetRequestId() *string
}

type ModifyHeadersResponseBody struct {
	// The unique ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// DC38A1D3-C042-5670-8394-8F6B1FA97B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHeadersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHeadersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHeadersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHeadersResponseBody) SetRequestId(v string) *ModifyHeadersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHeadersResponseBody) Validate() error {
	return dara.Validate(s)
}
