// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateLogConfigResponseBody
	GetRequestId() *string
}

type CreateLogConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLogConfigResponseBody) SetRequestId(v string) *CreateLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
