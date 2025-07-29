// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncClusterNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SyncClusterNodePoolResponseBody
	GetRequestId() *string
}

type SyncClusterNodePoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2D69A58F-345C-4FDE-88E4-BF51894XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncClusterNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *SyncClusterNodePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncClusterNodePoolResponseBody) SetRequestId(v string) *SyncClusterNodePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncClusterNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
