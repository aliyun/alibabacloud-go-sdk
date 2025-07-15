// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAudioAuditNotifyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveAudioAuditNotifyConfigResponseBody
	GetRequestId() *string
}

type AddLiveAudioAuditNotifyConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7BF95F2A-3B24-4CDE-9346-7F6FA86697A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveAudioAuditNotifyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAudioAuditNotifyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveAudioAuditNotifyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveAudioAuditNotifyConfigResponseBody) SetRequestId(v string) *AddLiveAudioAuditNotifyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveAudioAuditNotifyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
