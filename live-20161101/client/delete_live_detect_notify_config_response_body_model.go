// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveDetectNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveDetectNotifyConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveDetectNotifyConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveDetectNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveDetectNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveDetectNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveDetectNotifyConfigResponseBody) SetRequestId(v string) *DeleteLiveDetectNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveDetectNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
