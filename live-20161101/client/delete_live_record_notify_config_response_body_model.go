// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveRecordNotifyConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveRecordNotifyConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveRecordNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveRecordNotifyConfigResponseBody) SetRequestId(v string) *DeleteLiveRecordNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveRecordNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
