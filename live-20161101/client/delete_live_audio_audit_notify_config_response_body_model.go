// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAudioAuditNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveAudioAuditNotifyConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveAudioAuditNotifyConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7BF95F2A-3B24-4CDE-9346-7F6FA86697A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveAudioAuditNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAudioAuditNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveAudioAuditNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveAudioAuditNotifyConfigResponseBody) SetRequestId(v string) *DeleteLiveAudioAuditNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveAudioAuditNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
