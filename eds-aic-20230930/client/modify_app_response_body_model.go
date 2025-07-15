// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAppResponseBody
	GetRequestId() *string
}

type ModifyAppResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 83418504-5A82-5896-A24C-B2D468F0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppResponseBody) SetRequestId(v string) *ModifyAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppResponseBody) Validate() error {
	return dara.Validate(s)
}
