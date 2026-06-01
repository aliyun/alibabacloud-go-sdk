// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRayHistoryServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRayHistoryServerResponseBody
	GetRequestId() *string
}

type DeleteRayHistoryServerResponseBody struct {
	// example:
	//
	// 78F6FCE2-278F-4C4A-A6B7-DD8ECEA9C456
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRayHistoryServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRayHistoryServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRayHistoryServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRayHistoryServerResponseBody) SetRequestId(v string) *DeleteRayHistoryServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRayHistoryServerResponseBody) Validate() error {
	return dara.Validate(s)
}
