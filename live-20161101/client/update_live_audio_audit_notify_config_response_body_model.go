// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAudioAuditNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveAudioAuditNotifyConfigResponseBody
	GetRequestId() *string
}

type UpdateLiveAudioAuditNotifyConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7BF95F2A-3B24-4CDE-9346-7F6FA86697A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveAudioAuditNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAudioAuditNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveAudioAuditNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveAudioAuditNotifyConfigResponseBody) SetRequestId(v string) *UpdateLiveAudioAuditNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveAudioAuditNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
