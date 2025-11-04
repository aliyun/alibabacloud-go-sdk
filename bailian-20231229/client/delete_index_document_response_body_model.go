// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexDocumentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteIndexDocumentResponseBody
	GetCode() *string
	SetData(v *DeleteIndexDocumentResponseBodyData) *DeleteIndexDocumentResponseBody
	GetData() *DeleteIndexDocumentResponseBodyData
	SetMessage(v string) *DeleteIndexDocumentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIndexDocumentResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteIndexDocumentResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DeleteIndexDocumentResponseBody
	GetSuccess() *bool
}

type DeleteIndexDocumentResponseBody struct {
	// HTTP status code
	//
	// example:
	//
	// Index.InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The parameters returned by the operation.
	Data *DeleteIndexDocumentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Required parameter(%s) missing or invalid, please check the request parameters.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indications whether the API call is successful. Valid values:
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

func (s DeleteIndexDocumentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexDocumentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteIndexDocumentResponseBody) GetData() *DeleteIndexDocumentResponseBodyData {
	return s.Data
}

func (s *DeleteIndexDocumentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIndexDocumentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIndexDocumentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteIndexDocumentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteIndexDocumentResponseBody) SetCode(v string) *DeleteIndexDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIndexDocumentResponseBody) SetData(v *DeleteIndexDocumentResponseBodyData) *DeleteIndexDocumentResponseBody {
	s.Data = v
	return s
}

func (s *DeleteIndexDocumentResponseBody) SetMessage(v string) *DeleteIndexDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIndexDocumentResponseBody) SetRequestId(v string) *DeleteIndexDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexDocumentResponseBody) SetStatus(v string) *DeleteIndexDocumentResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteIndexDocumentResponseBody) SetSuccess(v bool) *DeleteIndexDocumentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIndexDocumentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteIndexDocumentResponseBodyData struct {
	// The list of primary key IDs of documents that are deleted.
	DeletedDocument []*string `json:"DeletedDocument,omitempty" xml:"DeletedDocument,omitempty" type:"Repeated"`
}

func (s DeleteIndexDocumentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexDocumentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteIndexDocumentResponseBodyData) GetDeletedDocument() []*string {
	return s.DeletedDocument
}

func (s *DeleteIndexDocumentResponseBodyData) SetDeletedDocument(v []*string) *DeleteIndexDocumentResponseBodyData {
	s.DeletedDocument = v
	return s
}

func (s *DeleteIndexDocumentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
