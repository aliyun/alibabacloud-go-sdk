// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUpdateTaskResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *GetUpdateTaskResultRequest
	GetOperationId() *string
}

type GetUpdateTaskResultRequest struct {
	// The operation ID, which is used to query the result of the asynchronous node update. You can obtain this ID from the UpdateTaskAsync operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// e15ad21c-b0e9-4792-8f55-b037xxxxxxxx
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s GetUpdateTaskResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUpdateTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetUpdateTaskResultRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *GetUpdateTaskResultRequest) SetOperationId(v string) *GetUpdateTaskResultRequest {
	s.OperationId = &v
	return s
}

func (s *GetUpdateTaskResultRequest) Validate() error {
	return dara.Validate(s)
}
