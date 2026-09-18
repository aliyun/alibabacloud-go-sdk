// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyConnectorResponseBody
	GetCode() *string
	SetData(v *VerifyConnectorResponseBodyData) *VerifyConnectorResponseBody
	GetData() *VerifyConnectorResponseBodyData
	SetHttpStatusCode(v int32) *VerifyConnectorResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *VerifyConnectorResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyConnectorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyConnectorResponseBody
	GetSuccess() *bool
}

type VerifyConnectorResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The validation result.
	Data *VerifyConnectorResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s VerifyConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyConnectorResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyConnectorResponseBody) GetData() *VerifyConnectorResponseBodyData {
	return s.Data
}

func (s *VerifyConnectorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *VerifyConnectorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyConnectorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyConnectorResponseBody) SetCode(v string) *VerifyConnectorResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyConnectorResponseBody) SetData(v *VerifyConnectorResponseBodyData) *VerifyConnectorResponseBody {
	s.Data = v
	return s
}

func (s *VerifyConnectorResponseBody) SetHttpStatusCode(v int32) *VerifyConnectorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *VerifyConnectorResponseBody) SetMessage(v string) *VerifyConnectorResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyConnectorResponseBody) SetRequestId(v string) *VerifyConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyConnectorResponseBody) SetSuccess(v bool) *VerifyConnectorResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyConnectorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyConnectorResponseBodyData struct {
	// The list of Service Account Key labels that failed validation or returned indeterminate results.
	//
	// example:
	//
	// default
	InvalidServiceAccountKeys []*string `json:"invalidServiceAccountKeys,omitempty" xml:"invalidServiceAccountKeys,omitempty" type:"Repeated"`
	// Indicates whether all validated Service Account Keys are valid.
	//
	// example:
	//
	// true
	Valid *bool `json:"valid,omitempty" xml:"valid,omitempty"`
}

func (s VerifyConnectorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VerifyConnectorResponseBodyData) GoString() string {
	return s.String()
}

func (s *VerifyConnectorResponseBodyData) GetInvalidServiceAccountKeys() []*string {
	return s.InvalidServiceAccountKeys
}

func (s *VerifyConnectorResponseBodyData) GetValid() *bool {
	return s.Valid
}

func (s *VerifyConnectorResponseBodyData) SetInvalidServiceAccountKeys(v []*string) *VerifyConnectorResponseBodyData {
	s.InvalidServiceAccountKeys = v
	return s
}

func (s *VerifyConnectorResponseBodyData) SetValid(v bool) *VerifyConnectorResponseBodyData {
	s.Valid = &v
	return s
}

func (s *VerifyConnectorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
