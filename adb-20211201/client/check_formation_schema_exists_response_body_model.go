// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFormationSchemaExistsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckFormationSchemaExistsResponseBody
	GetCode() *string
	SetData(v string) *CheckFormationSchemaExistsResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CheckFormationSchemaExistsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CheckFormationSchemaExistsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckFormationSchemaExistsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckFormationSchemaExistsResponseBody
	GetSuccess() *bool
}

type CheckFormationSchemaExistsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the schema exists.
	//
	// example:
	//
	// false
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message. OK is returned if the request was successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 82C80833-BE26-424D-B485-1ABC8431E14
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckFormationSchemaExistsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckFormationSchemaExistsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckFormationSchemaExistsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckFormationSchemaExistsResponseBody) GetData() *string {
	return s.Data
}

func (s *CheckFormationSchemaExistsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckFormationSchemaExistsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckFormationSchemaExistsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckFormationSchemaExistsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckFormationSchemaExistsResponseBody) SetCode(v string) *CheckFormationSchemaExistsResponseBody {
	s.Code = &v
	return s
}

func (s *CheckFormationSchemaExistsResponseBody) SetData(v string) *CheckFormationSchemaExistsResponseBody {
	s.Data = &v
	return s
}

func (s *CheckFormationSchemaExistsResponseBody) SetHttpStatusCode(v int32) *CheckFormationSchemaExistsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckFormationSchemaExistsResponseBody) SetMessage(v string) *CheckFormationSchemaExistsResponseBody {
	s.Message = &v
	return s
}

func (s *CheckFormationSchemaExistsResponseBody) SetRequestId(v string) *CheckFormationSchemaExistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckFormationSchemaExistsResponseBody) SetSuccess(v bool) *CheckFormationSchemaExistsResponseBody {
	s.Success = &v
	return s
}

func (s *CheckFormationSchemaExistsResponseBody) Validate() error {
	return dara.Validate(s)
}
