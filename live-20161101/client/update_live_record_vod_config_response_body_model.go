// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveRecordVodConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveRecordVodConfigResponseBody
	GetRequestId() *string
}

type UpdateLiveRecordVodConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 61C96B9A-F203-4EC5-8E43-CB92E68F67DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveRecordVodConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveRecordVodConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveRecordVodConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveRecordVodConfigResponseBody) SetRequestId(v string) *UpdateLiveRecordVodConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveRecordVodConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
