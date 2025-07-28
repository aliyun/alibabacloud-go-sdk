// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefaultHttpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDefaultHttpsResponseBody
	GetRequestId() *string
}

type ModifyDefaultHttpsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 276D7566-31C9-4192-9DD1-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDefaultHttpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefaultHttpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDefaultHttpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDefaultHttpsResponseBody) SetRequestId(v string) *ModifyDefaultHttpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDefaultHttpsResponseBody) Validate() error {
	return dara.Validate(s)
}
