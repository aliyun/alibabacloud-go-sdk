// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVsPullStreamInfoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVsPullStreamInfoConfigResponseBody
	GetRequestId() *string
}

type DeleteVsPullStreamInfoConfigResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVsPullStreamInfoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVsPullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVsPullStreamInfoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVsPullStreamInfoConfigResponseBody) SetRequestId(v string) *DeleteVsPullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVsPullStreamInfoConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
