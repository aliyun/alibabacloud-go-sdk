// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateDNSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePrivateDNSResponseBody
	GetRequestId() *string
}

type CreatePrivateDNSResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePrivateDNSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateDNSResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrivateDNSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrivateDNSResponseBody) SetRequestId(v string) *CreatePrivateDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrivateDNSResponseBody) Validate() error {
	return dara.Validate(s)
}
