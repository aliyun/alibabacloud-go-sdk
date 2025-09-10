// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggTaskGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAggTaskGroupResponseBody
	GetRequestId() *string
}

type DeleteAggTaskGroupResponseBody struct {
	// example:
	//
	// 264C3E89-XXXX-XXXX-XXXX-CE9C2196C7DC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAggTaskGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAggTaskGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAggTaskGroupResponseBody) SetRequestId(v string) *DeleteAggTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAggTaskGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
