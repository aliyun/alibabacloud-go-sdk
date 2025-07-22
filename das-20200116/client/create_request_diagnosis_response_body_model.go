// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRequestDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRequestDiagnosisResponseBody
	GetCode() *string
	SetData(v string) *CreateRequestDiagnosisResponseBody
	GetData() *string
	SetMessage(v string) *CreateRequestDiagnosisResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRequestDiagnosisResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateRequestDiagnosisResponseBody
	GetSuccess() *string
}

type CreateRequestDiagnosisResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The diagnostics ID, which is the unique identifier of the diagnosis. This parameter can be used to query the result of the diagnosis.
	//
	// example:
	//
	// 61820b594664275c4429****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 800FBAF5-A539-5B97-A09E-C63AB2F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRequestDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRequestDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRequestDiagnosisResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRequestDiagnosisResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateRequestDiagnosisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRequestDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRequestDiagnosisResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateRequestDiagnosisResponseBody) SetCode(v string) *CreateRequestDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRequestDiagnosisResponseBody) SetData(v string) *CreateRequestDiagnosisResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRequestDiagnosisResponseBody) SetMessage(v string) *CreateRequestDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRequestDiagnosisResponseBody) SetRequestId(v string) *CreateRequestDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRequestDiagnosisResponseBody) SetSuccess(v string) *CreateRequestDiagnosisResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRequestDiagnosisResponseBody) Validate() error {
	return dara.Validate(s)
}
