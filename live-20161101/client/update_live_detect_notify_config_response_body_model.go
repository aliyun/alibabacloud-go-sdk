// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveDetectNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveDetectNotifyConfigResponseBody
	GetRequestId() *string
}

type UpdateLiveDetectNotifyConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveDetectNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveDetectNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveDetectNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveDetectNotifyConfigResponseBody) SetRequestId(v string) *UpdateLiveDetectNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveDetectNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
