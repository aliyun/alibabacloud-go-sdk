// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendLiveSnapshotJobCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendLiveSnapshotJobCommandResponseBody
	GetRequestId() *string
}

type SendLiveSnapshotJobCommandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendLiveSnapshotJobCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendLiveSnapshotJobCommandResponseBody) GoString() string {
	return s.String()
}

func (s *SendLiveSnapshotJobCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendLiveSnapshotJobCommandResponseBody) SetRequestId(v string) *SendLiveSnapshotJobCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendLiveSnapshotJobCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
