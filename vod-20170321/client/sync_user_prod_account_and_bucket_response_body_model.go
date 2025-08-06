// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncUserProdAccountAndBucketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SyncUserProdAccountAndBucketResponseBody
	GetRequestId() *string
	SetResult(v bool) *SyncUserProdAccountAndBucketResponseBody
	GetResult() *bool
}

type SyncUserProdAccountAndBucketResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SyncUserProdAccountAndBucketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncUserProdAccountAndBucketResponseBody) GoString() string {
	return s.String()
}

func (s *SyncUserProdAccountAndBucketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncUserProdAccountAndBucketResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SyncUserProdAccountAndBucketResponseBody) SetRequestId(v string) *SyncUserProdAccountAndBucketResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncUserProdAccountAndBucketResponseBody) SetResult(v bool) *SyncUserProdAccountAndBucketResponseBody {
	s.Result = &v
	return s
}

func (s *SyncUserProdAccountAndBucketResponseBody) Validate() error {
	return dara.Validate(s)
}
