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
	// The ID of the refresh task. Multiple IDs are separated by commas (,). Refresh tasks are merged based on the following rules:
	//
	// If the tasks are set for the same accelerated domain name, submitted within the same second, and refresh content based on URLs instead of directories, the tasks IDs are merged into the same task ID (RefreshTaskId). If the number of these tasks exceeds 2,000, every 2,000 tasks IDs are merged into the same task ID (RefreshTaskId).
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
