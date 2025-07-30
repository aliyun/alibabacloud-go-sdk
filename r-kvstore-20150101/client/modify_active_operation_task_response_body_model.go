// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActiveOperationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *ModifyActiveOperationTaskResponseBody
	GetIds() *string
	SetRequestId(v string) *ModifyActiveOperationTaskResponseBody
	GetRequestId() *string
}

type ModifyActiveOperationTaskResponseBody struct {
	// The ID of the O\\&M task. IDs are separated by commas (,).
	//
	// example:
	//
	// 11111,22222
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyActiveOperationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyActiveOperationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyActiveOperationTaskResponseBody) GetIds() *string {
	return s.Ids
}

func (s *ModifyActiveOperationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyActiveOperationTaskResponseBody) SetIds(v string) *ModifyActiveOperationTaskResponseBody {
	s.Ids = &v
	return s
}

func (s *ModifyActiveOperationTaskResponseBody) SetRequestId(v string) *ModifyActiveOperationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyActiveOperationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
