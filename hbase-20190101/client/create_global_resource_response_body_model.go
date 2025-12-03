// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateGlobalResourceResponseBody
	GetRequestId() *string
}

type CreateGlobalResourceResponseBody struct {
	// example:
	//
	// 1AB9ABDF-7E1E-44AD-8610-70A005115DD1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGlobalResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGlobalResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGlobalResourceResponseBody) SetRequestId(v string) *CreateGlobalResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGlobalResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
