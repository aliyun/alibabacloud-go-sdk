// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendSystemPropertyTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendSystemPropertyTemplateResponseBody
	GetRequestId() *string
}

type SendSystemPropertyTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendSystemPropertyTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendSystemPropertyTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *SendSystemPropertyTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendSystemPropertyTemplateResponseBody) SetRequestId(v string) *SendSystemPropertyTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendSystemPropertyTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
