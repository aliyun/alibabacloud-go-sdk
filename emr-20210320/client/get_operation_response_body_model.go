// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperation(v *Operation) *GetOperationResponseBody
	GetOperation() *Operation
	SetRequestId(v string) *GetOperationResponseBody
	GetRequestId() *string
}

type GetOperationResponseBody struct {
	// The operation that was performed.
	Operation *Operation `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA3896A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetOperationResponseBody) GetOperation() *Operation {
	return s.Operation
}

func (s *GetOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOperationResponseBody) SetOperation(v *Operation) *GetOperationResponseBody {
	s.Operation = v
	return s
}

func (s *GetOperationResponseBody) SetRequestId(v string) *GetOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOperationResponseBody) Validate() error {
	if s.Operation != nil {
		if err := s.Operation.Validate(); err != nil {
			return err
		}
	}
	return nil
}
