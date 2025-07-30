// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDtsStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SyncDtsStatusResponseBody
	GetRequestId() *string
}

type SyncDtsStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncDtsStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncDtsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDtsStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncDtsStatusResponseBody) SetRequestId(v string) *SyncDtsStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncDtsStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
