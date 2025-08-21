// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVsStreamsNotifyUrlConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVsStreamsNotifyUrlConfigResponseBody
	GetRequestId() *string
}

type DeleteVsStreamsNotifyUrlConfigResponseBody struct {
	// example:
	//
	// 4C747C97-7ECD-4C61-8A92-67AD806331FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVsStreamsNotifyUrlConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVsStreamsNotifyUrlConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVsStreamsNotifyUrlConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVsStreamsNotifyUrlConfigResponseBody) SetRequestId(v string) *DeleteVsStreamsNotifyUrlConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVsStreamsNotifyUrlConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
