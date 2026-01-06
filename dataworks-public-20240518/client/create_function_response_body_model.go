// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateFunctionResponseBody
	GetId() *string
	SetRequestId(v string) *CreateFunctionResponseBody
	GetRequestId() *string
}

type CreateFunctionResponseBody struct {
	// The unique identifier of the UDF function.
	//
	// >  Prior to SDK version 8.0.0, this field is of type Long. In SDK version 8.0.0 and later, it is of type String. This change does not affect the normal use of the SDK. The parameter is returned based on the type defined in the SDK. Compilation failures caused by the type change may occur only when you upgrade the SDK across version 8.0.0. In this case, you must manually update the data type.
	//
	// example:
	//
	// 580667964888595XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// AE49C88D-5BEE-5ADD-8B8C-C4BBC0D7XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFunctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFunctionResponseBody) SetId(v string) *CreateFunctionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFunctionResponseBody) SetRequestId(v string) *CreateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionResponseBody) Validate() error {
	return dara.Validate(s)
}
