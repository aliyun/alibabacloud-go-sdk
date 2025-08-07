// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDBClusterStorageUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefreshDBClusterStorageUsageResponseBody
	GetRequestId() *string
}

type RefreshDBClusterStorageUsageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshDBClusterStorageUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshDBClusterStorageUsageResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDBClusterStorageUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshDBClusterStorageUsageResponseBody) SetRequestId(v string) *RefreshDBClusterStorageUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDBClusterStorageUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
