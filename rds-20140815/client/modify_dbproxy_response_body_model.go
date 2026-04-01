// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBProxyResponseBody
	GetRequestId() *string
}

type ModifyDBProxyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FC452BB1-EED8-4278-95C7-0324B3710DF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBProxyResponseBody) SetRequestId(v string) *ModifyDBProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
