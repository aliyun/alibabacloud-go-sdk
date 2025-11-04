// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateIndexResponseBody
	GetCode() *string
	SetData(v *CreateIndexResponseBodyData) *CreateIndexResponseBody
	GetData() *CreateIndexResponseBodyData
	SetMessage(v string) *CreateIndexResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateIndexResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateIndexResponseBody
	GetStatus() *string
	SetSuccess(v bool) *CreateIndexResponseBody
	GetSuccess() *bool
}

type CreateIndexResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Forbidden
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned if the request is successful.
	Data *CreateIndexResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Invalid input, variable name is missing
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17204B98-7734-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndexResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateIndexResponseBody) GetData() *CreateIndexResponseBodyData {
	return s.Data
}

func (s *CreateIndexResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIndexResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateIndexResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateIndexResponseBody) SetCode(v string) *CreateIndexResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIndexResponseBody) SetData(v *CreateIndexResponseBodyData) *CreateIndexResponseBody {
	s.Data = v
	return s
}

func (s *CreateIndexResponseBody) SetMessage(v string) *CreateIndexResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIndexResponseBody) SetRequestId(v string) *CreateIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndexResponseBody) SetStatus(v string) *CreateIndexResponseBody {
	s.Status = &v
	return s
}

func (s *CreateIndexResponseBody) SetSuccess(v bool) *CreateIndexResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIndexResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIndexResponseBodyData struct {
	// The knowledge base ID, or `IndexId`, is a unique identifier for the knowledge base created.
	//
	// > Keep this ID. It is required for all subsequent API operations related to this knowledge base.
	//
	// example:
	//
	// jkurxhju6b
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateIndexResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIndexResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CreateIndexResponseBodyData) SetId(v string) *CreateIndexResponseBodyData {
	s.Id = &v
	return s
}

func (s *CreateIndexResponseBodyData) Validate() error {
	return dara.Validate(s)
}
