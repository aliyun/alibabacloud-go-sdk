// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceTrailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteServiceTrailResponseBody
	GetRequestId() *string
}

type DeleteServiceTrailResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20456DD5-5CBF-5015-9173-12CA4246B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceTrailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceTrailResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceTrailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceTrailResponseBody) SetRequestId(v string) *DeleteServiceTrailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceTrailResponseBody) Validate() error {
	return dara.Validate(s)
}
