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
	// Error status code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *CreateIndexResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 17204B98-xxxx-4F9A--2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned by the interface.
	//
	// example:
	//
	// "200"
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: Successful
	//
	// - false: Failed
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
	// The ID of the knowledge base, also known as \\`IndexId\\`. This is the unique identifier of the created knowledge base.
	//
	// > Keep this value safe. It will be used for all subsequent API operations related to this knowledge base.
	//
	// example:
	//
	// jkurxhxxxx
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
