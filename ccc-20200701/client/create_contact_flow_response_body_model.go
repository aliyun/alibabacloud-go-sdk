// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateContactFlowResponseBody
	GetCode() *string
	SetData(v string) *CreateContactFlowResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateContactFlowResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateContactFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateContactFlowResponseBody
	GetRequestId() *string
}

type CreateContactFlowResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 78128960-bb00-4ddc-8e82-923a8c5bd22d
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2778FA12-EDD6-42AA-9B15-AF855072E5E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateContactFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContactFlowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContactFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateContactFlowResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateContactFlowResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateContactFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateContactFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContactFlowResponseBody) SetCode(v string) *CreateContactFlowResponseBody {
	s.Code = &v
	return s
}

func (s *CreateContactFlowResponseBody) SetData(v string) *CreateContactFlowResponseBody {
	s.Data = &v
	return s
}

func (s *CreateContactFlowResponseBody) SetHttpStatusCode(v int32) *CreateContactFlowResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateContactFlowResponseBody) SetMessage(v string) *CreateContactFlowResponseBody {
	s.Message = &v
	return s
}

func (s *CreateContactFlowResponseBody) SetRequestId(v string) *CreateContactFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContactFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
