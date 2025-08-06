// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMessageCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetMessageCallbackResponseBody
	GetRequestId() *string
}

type SetMessageCallbackResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetMessageCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetMessageCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *SetMessageCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetMessageCallbackResponseBody) SetRequestId(v string) *SetMessageCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetMessageCallbackResponseBody) Validate() error {
	return dara.Validate(s)
}
