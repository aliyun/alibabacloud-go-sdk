// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateShardTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateShardTaskResponseBody
	GetData() *bool
	SetRequestId(v string) *CreateShardTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateShardTaskResponseBody
	GetSuccess() *bool
}

type CreateShardTaskResponseBody struct {
	// Task creation result
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F8997D95-94AD-416A-AE70-E24D08******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the operation.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateShardTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateShardTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateShardTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateShardTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateShardTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateShardTaskResponseBody) SetData(v bool) *CreateShardTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateShardTaskResponseBody) SetRequestId(v string) *CreateShardTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateShardTaskResponseBody) SetSuccess(v bool) *CreateShardTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateShardTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
