// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateGroupResponseBody
	GetRequestId() *string
}

type DissociateGroupResponseBody struct {
	// example:
	//
	// 17793D91-A26F-520D-A948-3452A45D15B1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DissociateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateGroupResponseBody) SetRequestId(v string) *DissociateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
