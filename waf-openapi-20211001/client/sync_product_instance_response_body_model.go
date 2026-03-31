// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncProductInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SyncProductInstanceResponseBody
	GetRequestId() *string
}

type SyncProductInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 45BA2382-7C3F-5B29-9A83-C3BCE586****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncProductInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncProductInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *SyncProductInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncProductInstanceResponseBody) SetRequestId(v string) *SyncProductInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncProductInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
