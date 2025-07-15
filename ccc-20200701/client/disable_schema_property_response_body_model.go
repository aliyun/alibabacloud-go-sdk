// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSchemaPropertyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableSchemaPropertyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DisableSchemaPropertyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DisableSchemaPropertyResponseBody
	GetMessage() *string
	SetParams(v []*string) *DisableSchemaPropertyResponseBody
	GetParams() []*string
	SetRequestId(v string) *DisableSchemaPropertyResponseBody
	GetRequestId() *string
}

type DisableSchemaPropertyResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// User 244715989906081477 does not exist in instance worldfirst01.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 9FBA26B0-462B-4D77-B78F-AF35560DBC71
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableSchemaPropertyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableSchemaPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSchemaPropertyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableSchemaPropertyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DisableSchemaPropertyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableSchemaPropertyResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DisableSchemaPropertyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableSchemaPropertyResponseBody) SetCode(v string) *DisableSchemaPropertyResponseBody {
	s.Code = &v
	return s
}

func (s *DisableSchemaPropertyResponseBody) SetHttpStatusCode(v int32) *DisableSchemaPropertyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DisableSchemaPropertyResponseBody) SetMessage(v string) *DisableSchemaPropertyResponseBody {
	s.Message = &v
	return s
}

func (s *DisableSchemaPropertyResponseBody) SetParams(v []*string) *DisableSchemaPropertyResponseBody {
	s.Params = v
	return s
}

func (s *DisableSchemaPropertyResponseBody) SetRequestId(v string) *DisableSchemaPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSchemaPropertyResponseBody) Validate() error {
	return dara.Validate(s)
}
