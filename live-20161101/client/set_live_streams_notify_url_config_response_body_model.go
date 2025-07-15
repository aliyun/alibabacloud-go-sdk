// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamsNotifyUrlConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLiveStreamsNotifyUrlConfigResponseBody
	GetRequestId() *string
}

type SetLiveStreamsNotifyUrlConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 40A4F36D-A7CC-473A-88E7-154F92242566
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLiveStreamsNotifyUrlConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamsNotifyUrlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveStreamsNotifyUrlConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveStreamsNotifyUrlConfigResponseBody) SetRequestId(v string) *SetLiveStreamsNotifyUrlConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveStreamsNotifyUrlConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
