// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexJobStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIndexJobStatusResponseBody
	GetCode() *string
	SetData(v *GetIndexJobStatusResponseBodyData) *GetIndexJobStatusResponseBody
	GetData() *GetIndexJobStatusResponseBodyData
	SetMessage(v string) *GetIndexJobStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIndexJobStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetIndexJobStatusResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetIndexJobStatusResponseBody
	GetSuccess() *bool
}

type GetIndexJobStatusResponseBody struct {
	// HTTP status code
	//
	// example:
	//
	// Index.Forbidden
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetIndexJobStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code returned.
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

func (s GetIndexJobStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIndexJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIndexJobStatusResponseBody) GetData() *GetIndexJobStatusResponseBodyData {
	return s.Data
}

func (s *GetIndexJobStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIndexJobStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIndexJobStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetIndexJobStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIndexJobStatusResponseBody) SetCode(v string) *GetIndexJobStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetData(v *GetIndexJobStatusResponseBodyData) *GetIndexJobStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetMessage(v string) *GetIndexJobStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetRequestId(v string) *GetIndexJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetStatus(v string) *GetIndexJobStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) SetSuccess(v bool) *GetIndexJobStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetIndexJobStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIndexJobStatusResponseBodyData struct {
	// The list of imported documents.
	Documents []*GetIndexJobStatusResponseBodyDataDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// The ID of the job.
	//
	// example:
	//
	// 66122af12a4e45ddae6bd6c845556647
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The status of the knowledge base job. Valid values:
	//
	// 	- COMPLETED
	//
	// 	- FAILED
	//
	// 	- RUNNING
	//
	// 	- PENDING
	//
	// example:
	//
	// PENDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetIndexJobStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIndexJobStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponseBodyData) GetDocuments() []*GetIndexJobStatusResponseBodyDataDocuments {
	return s.Documents
}

func (s *GetIndexJobStatusResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *GetIndexJobStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetIndexJobStatusResponseBodyData) SetDocuments(v []*GetIndexJobStatusResponseBodyDataDocuments) *GetIndexJobStatusResponseBodyData {
	s.Documents = v
	return s
}

func (s *GetIndexJobStatusResponseBodyData) SetJobId(v string) *GetIndexJobStatusResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyData) SetStatus(v string) *GetIndexJobStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyData) Validate() error {
	if s.Documents != nil {
		for _, item := range s.Documents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetIndexJobStatusResponseBodyDataDocuments struct {
	// HTTP status code
	//
	// example:
	//
	// Index.Document.ChunkError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The primary key ID of the document.
	//
	// example:
	//
	// file_9a65732555b54d5ea10796ca5742ba22_XXXXXXXX
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// The name of the document.
	DocName     *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The error message.
	//
	// example:
	//
	// document parse error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The import status of the document. Valid values:
	//
	// 	- INSERT_ERROR
	//
	// 	- RUNNING
	//
	// 	- DELETED
	//
	// 	- FINISH
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetIndexJobStatusResponseBodyDataDocuments) String() string {
	return dara.Prettify(s)
}

func (s GetIndexJobStatusResponseBodyDataDocuments) GoString() string {
	return s.String()
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) GetCode() *string {
	return s.Code
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) GetDocId() *string {
	return s.DocId
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) GetDocName() *string {
	return s.DocName
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) GetMessage() *string {
	return s.Message
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) GetStatus() *string {
	return s.Status
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetCode(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.Code = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetDocId(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.DocId = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetDocName(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.DocName = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetGmtModified(v int64) *GetIndexJobStatusResponseBodyDataDocuments {
	s.GmtModified = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetMessage(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.Message = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) SetStatus(v string) *GetIndexJobStatusResponseBodyDataDocuments {
	s.Status = &v
	return s
}

func (s *GetIndexJobStatusResponseBodyDataDocuments) Validate() error {
	return dara.Validate(s)
}
