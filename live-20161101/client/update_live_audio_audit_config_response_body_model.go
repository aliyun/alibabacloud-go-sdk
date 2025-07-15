// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAudioAuditConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveAudioAuditConfigResponseBody
	GetRequestId() *string
}

type UpdateLiveAudioAuditConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7BF95F2A-3B24-4CDE-9346-7F6FA86697A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveAudioAuditConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAudioAuditConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveAudioAuditConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveAudioAuditConfigResponseBody) SetRequestId(v string) *UpdateLiveAudioAuditConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
