// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstanceSnapshotResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteInstanceSnapshotResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *DeleteInstanceSnapshotResponseBody
	GetInstanceId() *string
	SetMessage(v string) *DeleteInstanceSnapshotResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstanceSnapshotResponseBody
	GetRequestId() *string
	SetSnapshotId(v string) *DeleteInstanceSnapshotResponseBody
	GetSnapshotId() *string
	SetSuccess(v bool) *DeleteInstanceSnapshotResponseBody
	GetSuccess() *bool
}

type DeleteInstanceSnapshotResponseBody struct {
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// null
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// snp-05cexxxxxxxxx
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstanceSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceSnapshotResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstanceSnapshotResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteInstanceSnapshotResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceSnapshotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstanceSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceSnapshotResponseBody) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DeleteInstanceSnapshotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstanceSnapshotResponseBody) SetCode(v string) *DeleteInstanceSnapshotResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetHttpStatusCode(v int32) *DeleteInstanceSnapshotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetInstanceId(v string) *DeleteInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetMessage(v string) *DeleteInstanceSnapshotResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetRequestId(v string) *DeleteInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetSnapshotId(v string) *DeleteInstanceSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) SetSuccess(v bool) *DeleteInstanceSnapshotResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstanceSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
