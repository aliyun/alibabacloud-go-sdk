// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveRecordNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveRecordNotifyConfigResponseBody
	GetRequestId() *string
}

type UpdateLiveRecordNotifyConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveRecordNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveRecordNotifyConfigResponseBody) SetRequestId(v string) *UpdateLiveRecordNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveRecordNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
