// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkerResourceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *UpdateWorkerResourceStatusResponseBody
	GetData() interface{}
	SetRequestId(v string) *UpdateWorkerResourceStatusResponseBody
	GetRequestId() *string
}

type UpdateWorkerResourceStatusResponseBody struct {
	// The response parameters.
	//
	// example:
	//
	// []
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWorkerResourceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResourceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResourceStatusResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateWorkerResourceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkerResourceStatusResponseBody) SetData(v interface{}) *UpdateWorkerResourceStatusResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWorkerResourceStatusResponseBody) SetRequestId(v string) *UpdateWorkerResourceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkerResourceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
