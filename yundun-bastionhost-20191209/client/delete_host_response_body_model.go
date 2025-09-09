// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHostResponseBody
	GetRequestId() *string
}

type DeleteHostResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHostResponseBody) SetRequestId(v string) *DeleteHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHostResponseBody) Validate() error {
	return dara.Validate(s)
}
