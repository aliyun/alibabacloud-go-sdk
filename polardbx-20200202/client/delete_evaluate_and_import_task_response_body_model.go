// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEvaluateAndImportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteEvaluateAndImportTaskResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteEvaluateAndImportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEvaluateAndImportTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEvaluateAndImportTaskResponseBody
	GetSuccess() *bool
}

type DeleteEvaluateAndImportTaskResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// msg
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 14036EBE-***-44DB-ACE9-AC6157D3A6FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEvaluateAndImportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEvaluateAndImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEvaluateAndImportTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteEvaluateAndImportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEvaluateAndImportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEvaluateAndImportTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEvaluateAndImportTaskResponseBody) SetData(v bool) *DeleteEvaluateAndImportTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteEvaluateAndImportTaskResponseBody) SetMessage(v string) *DeleteEvaluateAndImportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEvaluateAndImportTaskResponseBody) SetRequestId(v string) *DeleteEvaluateAndImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEvaluateAndImportTaskResponseBody) SetSuccess(v bool) *DeleteEvaluateAndImportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEvaluateAndImportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
