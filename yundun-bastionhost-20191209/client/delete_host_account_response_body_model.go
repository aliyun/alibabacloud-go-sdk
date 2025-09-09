// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHostAccountResponseBody
	GetRequestId() *string
}

type DeleteHostAccountResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHostAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHostAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHostAccountResponseBody) SetRequestId(v string) *DeleteHostAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHostAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
