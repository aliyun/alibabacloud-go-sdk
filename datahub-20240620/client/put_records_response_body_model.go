// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutRecordsResponseBody
	GetRequestId() *string
	SetShardId(v string) *PutRecordsResponseBody
	GetShardId() *string
	SetSuccess(v bool) *PutRecordsResponseBody
	GetSuccess() *bool
}

type PutRecordsResponseBody struct {
	// example:
	//
	// A20A7093-8FE0-058C-BE0C-3C8057D5F1A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 7
	ShardId *string `json:"ShardId,omitempty" xml:"ShardId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *PutRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutRecordsResponseBody) GetShardId() *string {
	return s.ShardId
}

func (s *PutRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutRecordsResponseBody) SetRequestId(v string) *PutRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutRecordsResponseBody) SetShardId(v string) *PutRecordsResponseBody {
	s.ShardId = &v
	return s
}

func (s *PutRecordsResponseBody) SetSuccess(v bool) *PutRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *PutRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}
