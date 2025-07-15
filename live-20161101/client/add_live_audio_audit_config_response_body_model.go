// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAudioAuditConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveAudioAuditConfigResponseBody
	GetRequestId() *string
}

type AddLiveAudioAuditConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7BF95F2A-3B24-4CDE-9346-7F6FA86697A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveAudioAuditConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAudioAuditConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveAudioAuditConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveAudioAuditConfigResponseBody) SetRequestId(v string) *AddLiveAudioAuditConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveAudioAuditConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
