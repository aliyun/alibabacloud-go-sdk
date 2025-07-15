// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamsNotifyUrlConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveStreamsNotifyUrlConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveStreamsNotifyUrlConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 40A4F36D-A7CC-473A-88E7-154F92242566
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveStreamsNotifyUrlConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamsNotifyUrlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamsNotifyUrlConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveStreamsNotifyUrlConfigResponseBody) SetRequestId(v string) *DeleteLiveStreamsNotifyUrlConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveStreamsNotifyUrlConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
