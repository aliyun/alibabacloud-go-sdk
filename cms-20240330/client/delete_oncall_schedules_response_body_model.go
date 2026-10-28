// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOncallSchedulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteOncallSchedulesResponseBody
	GetRequestId() *string
}

type DeleteOncallSchedulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteOncallSchedulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOncallSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOncallSchedulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOncallSchedulesResponseBody) SetRequestId(v string) *DeleteOncallSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOncallSchedulesResponseBody) Validate() error {
	return dara.Validate(s)
}
