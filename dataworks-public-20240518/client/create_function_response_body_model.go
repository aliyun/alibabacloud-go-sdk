// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFunctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateFunctionResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateFunctionResponseBody
	GetRequestId() *string
}

type CreateFunctionResponseBody struct {
	// The ID of the UDF.
	//
	// example:
	//
	// 580667964888595XXXX
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *CreateFunctionResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateFunctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFunctionResponseBody) SetId(v int64) *CreateFunctionResponseBody {
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
