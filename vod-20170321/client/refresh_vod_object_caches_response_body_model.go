// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshVodObjectCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRefreshTaskId(v string) *RefreshVodObjectCachesResponseBody
	GetRefreshTaskId() *string
	SetRequestId(v string) *RefreshVodObjectCachesResponseBody
	GetRequestId() *string
}

type RefreshVodObjectCachesResponseBody struct {
	// The ID of the purge task. Multiple task IDs are separated by commas (,).
	//
	// The returned purge task IDs are merged based on the following rules:
	//
	// Purge tasks (at URL granularity) submitted for the same domain name within the same second are merged into a single RefreshTaskId.
	//
	// If purge tasks (at URL granularity) submitted for the same domain name within the same second exceed 2,000, they are merged into one RefreshTaskId per 2,000 tasks.
	//
	// example:
	//
	// 70422*****2904
	RefreshTaskId *string `json:"RefreshTaskId,omitempty" xml:"RefreshTaskId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D61E4801-EAFF-4A63-****-FBF6CE1CFD1C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshVodObjectCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshVodObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshVodObjectCachesResponseBody) GetRefreshTaskId() *string {
	return s.RefreshTaskId
}

func (s *RefreshVodObjectCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshVodObjectCachesResponseBody) SetRefreshTaskId(v string) *RefreshVodObjectCachesResponseBody {
	s.RefreshTaskId = &v
	return s
}

func (s *RefreshVodObjectCachesResponseBody) SetRequestId(v string) *RefreshVodObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshVodObjectCachesResponseBody) Validate() error {
	return dara.Validate(s)
}
