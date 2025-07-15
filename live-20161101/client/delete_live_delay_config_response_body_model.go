// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDelayConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveDelayConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveDelayConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveDelayConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDelayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveDelayConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveDelayConfigResponseBody) SetRequestId(v string) *DeleteLiveDelayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveDelayConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
