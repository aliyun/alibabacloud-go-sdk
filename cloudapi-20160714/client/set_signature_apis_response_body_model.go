// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSignatureApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSignatureApisResponseBody
	GetRequestId() *string
}

type SetSignatureApisResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ004
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSignatureApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSignatureApisResponseBody) GoString() string {
	return s.String()
}

func (s *SetSignatureApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSignatureApisResponseBody) SetRequestId(v string) *SetSignatureApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSignatureApisResponseBody) Validate() error {
	return dara.Validate(s)
}
