// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetListenerAccessControlStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetListenerAccessControlStatusResponseBody
	GetRequestId() *string
}

type SetListenerAccessControlStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetListenerAccessControlStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetListenerAccessControlStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetListenerAccessControlStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetListenerAccessControlStatusResponseBody) SetRequestId(v string) *SetListenerAccessControlStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetListenerAccessControlStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
