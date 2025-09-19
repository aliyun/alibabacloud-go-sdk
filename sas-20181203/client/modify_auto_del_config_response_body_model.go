// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoDelConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAutoDelConfigResponseBody
	GetRequestId() *string
}

type ModifyAutoDelConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3BFA659D-F44F-5703-8FD1-33AB596BEACA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAutoDelConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoDelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoDelConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAutoDelConfigResponseBody) SetRequestId(v string) *ModifyAutoDelConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAutoDelConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
