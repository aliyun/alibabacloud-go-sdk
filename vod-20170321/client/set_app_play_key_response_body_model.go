// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAppPlayKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetAppPlayKeyResponseBody
	GetRequestId() *string
}

type SetAppPlayKeyResponseBody struct {
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAppPlayKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAppPlayKeyResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppPlayKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAppPlayKeyResponseBody) SetRequestId(v string) *SetAppPlayKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppPlayKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
