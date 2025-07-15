// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAudioAuditConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveAudioAuditConfigResponseBody
	GetRequestId() *string
}

type DeleteLiveAudioAuditConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveAudioAuditConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAudioAuditConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveAudioAuditConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveAudioAuditConfigResponseBody) SetRequestId(v string) *DeleteLiveAudioAuditConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveAudioAuditConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
