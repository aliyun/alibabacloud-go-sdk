// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceSnapshotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateInstanceSnapshotResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateInstanceSnapshotResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *CreateInstanceSnapshotResponseBody
	GetInstanceId() *string
	SetMessage(v string) *CreateInstanceSnapshotResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateInstanceSnapshotResponseBody
	GetRequestId() *string
	SetSnapshotId(v string) *CreateInstanceSnapshotResponseBody
	GetSnapshotId() *string
	SetSuccess(v bool) *CreateInstanceSnapshotResponseBody
	GetSuccess() *bool
}

type CreateInstanceSnapshotResponseBody struct {
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

func (s CreateInstanceSnapshotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceSnapshotResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateInstanceSnapshotResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateInstanceSnapshotResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceSnapshotResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateInstanceSnapshotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceSnapshotResponseBody) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateInstanceSnapshotResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceSnapshotResponseBody) SetCode(v string) *CreateInstanceSnapshotResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetHttpStatusCode(v int32) *CreateInstanceSnapshotResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetInstanceId(v string) *CreateInstanceSnapshotResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetMessage(v string) *CreateInstanceSnapshotResponseBody {
	s.Message = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetRequestId(v string) *CreateInstanceSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetSnapshotId(v string) *CreateInstanceSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) SetSuccess(v bool) *CreateInstanceSnapshotResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceSnapshotResponseBody) Validate() error {
	return dara.Validate(s)
}
