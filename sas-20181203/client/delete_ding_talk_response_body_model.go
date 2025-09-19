// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDingTalkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDingTalkResponseBody
	GetRequestId() *string
}

type DeleteDingTalkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BF3D4ACB-CE17-559F-B850-490E42CDDC7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDingTalkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDingTalkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDingTalkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDingTalkResponseBody) SetRequestId(v string) *DeleteDingTalkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDingTalkResponseBody) Validate() error {
	return dara.Validate(s)
}
