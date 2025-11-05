// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskReplicaPairProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProgress(v int32) *DescribeDiskReplicaPairProgressResponseBody
	GetProgress() *int32
	SetRecoverPoint(v int64) *DescribeDiskReplicaPairProgressResponseBody
	GetRecoverPoint() *int64
	SetRequestId(v string) *DescribeDiskReplicaPairProgressResponseBody
	GetRequestId() *string
}

type DescribeDiskReplicaPairProgressResponseBody struct {
	// The replication progress of the replication pair.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The timestamp that indicates the last recovery point in time. The value is returned only after the replication pair works for replicating data.
	//
	// example:
	//
	// 1661917424
	RecoverPoint *int64 `json:"RecoverPoint,omitempty" xml:"RecoverPoint,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAA478A0-BEE6-1D42-BEB6-A9CFEAD6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiskReplicaPairProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskReplicaPairProgressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiskReplicaPairProgressResponseBody) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeDiskReplicaPairProgressResponseBody) GetRecoverPoint() *int64 {
	return s.RecoverPoint
}

func (s *DescribeDiskReplicaPairProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiskReplicaPairProgressResponseBody) SetProgress(v int32) *DescribeDiskReplicaPairProgressResponseBody {
	s.Progress = &v
	return s
}

func (s *DescribeDiskReplicaPairProgressResponseBody) SetRecoverPoint(v int64) *DescribeDiskReplicaPairProgressResponseBody {
	s.RecoverPoint = &v
	return s
}

func (s *DescribeDiskReplicaPairProgressResponseBody) SetRequestId(v string) *DescribeDiskReplicaPairProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiskReplicaPairProgressResponseBody) Validate() error {
	return dara.Validate(s)
}
