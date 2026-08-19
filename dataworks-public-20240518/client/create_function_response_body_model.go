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
	// 	Notice: This field was of the Long type in SDK versions earlier than 8.0.0 and is of the String type in SDK 8.0.0 and later. **This change does not affect normal SDK usage, and the parameter is still returned in the type defined in the SDK**. Only when you upgrade across SDK version 8.0.0, the type change may cause project compilation failures, and you need to manually correct the data type.
	//
	// example:
	//
	// 580667964888595XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// AE49C88D-5BEE-5ADD-8B8C-C4BBC0D7****
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
