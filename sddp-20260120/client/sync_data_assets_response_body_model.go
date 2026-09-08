// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDataAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SyncDataAssetsResponseBody
	GetRequestId() *string
}

type SyncDataAssetsResponseBody struct {
	// example:
	//
	// 7C6D8E9F-1234-5678-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncDataAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncDataAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDataAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncDataAssetsResponseBody) SetRequestId(v string) *SyncDataAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncDataAssetsResponseBody) Validate() error {
	return dara.Validate(s)
}
