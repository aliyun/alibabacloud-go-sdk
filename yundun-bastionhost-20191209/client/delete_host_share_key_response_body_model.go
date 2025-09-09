// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostShareKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHostShareKeyResponseBody
	GetRequestId() *string
}

type DeleteHostShareKeyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHostShareKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostShareKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHostShareKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHostShareKeyResponseBody) SetRequestId(v string) *DeleteHostShareKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHostShareKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
