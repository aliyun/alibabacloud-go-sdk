// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpApiOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHttpApiOperationResponseBody
	GetCode() *string
	SetData(v *CreateHttpApiOperationResponseBodyData) *CreateHttpApiOperationResponseBody
	GetData() *CreateHttpApiOperationResponseBodyData
	SetMessage(v string) *CreateHttpApiOperationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHttpApiOperationResponseBody
	GetRequestId() *string
}

type CreateHttpApiOperationResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Operation information.
	Data *CreateHttpApiOperationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateHttpApiOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiOperationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpApiOperationResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHttpApiOperationResponseBody) GetData() *CreateHttpApiOperationResponseBodyData {
	return s.Data
}

func (s *CreateHttpApiOperationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHttpApiOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHttpApiOperationResponseBody) SetCode(v string) *CreateHttpApiOperationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHttpApiOperationResponseBody) SetData(v *CreateHttpApiOperationResponseBodyData) *CreateHttpApiOperationResponseBody {
	s.Data = v
	return s
}

func (s *CreateHttpApiOperationResponseBody) SetMessage(v string) *CreateHttpApiOperationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHttpApiOperationResponseBody) SetRequestId(v string) *CreateHttpApiOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHttpApiOperationResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateHttpApiOperationResponseBodyData struct {
	// Operation information.
	Operations []*CreateHttpApiOperationResponseBodyDataOperations `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
}

func (s CreateHttpApiOperationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiOperationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateHttpApiOperationResponseBodyData) GetOperations() []*CreateHttpApiOperationResponseBodyDataOperations {
	return s.Operations
}

func (s *CreateHttpApiOperationResponseBodyData) SetOperations(v []*CreateHttpApiOperationResponseBodyDataOperations) *CreateHttpApiOperationResponseBodyData {
	s.Operations = v
	return s
}

func (s *CreateHttpApiOperationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type CreateHttpApiOperationResponseBodyDataOperations struct {
	// Operation ID.
	//
	// example:
	//
	// op-xxx
	OperationId *string `json:"operationId,omitempty" xml:"operationId,omitempty"`
}

func (s CreateHttpApiOperationResponseBodyDataOperations) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpApiOperationResponseBodyDataOperations) GoString() string {
	return s.String()
}

func (s *CreateHttpApiOperationResponseBodyDataOperations) GetOperationId() *string {
	return s.OperationId
}

func (s *CreateHttpApiOperationResponseBodyDataOperations) SetOperationId(v string) *CreateHttpApiOperationResponseBodyDataOperations {
	s.OperationId = &v
	return s
}

func (s *CreateHttpApiOperationResponseBodyDataOperations) Validate() error {
	return dara.Validate(s)
}
