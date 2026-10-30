// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForkKnowledgeFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ForkKnowledgeFileResponseBodyData) *ForkKnowledgeFileResponseBody
	GetData() *ForkKnowledgeFileResponseBodyData
	SetRequestId(v string) *ForkKnowledgeFileResponseBody
	GetRequestId() *string
}

type ForkKnowledgeFileResponseBody struct {
	// The returned data.
	Data *ForkKnowledgeFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ForkKnowledgeFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ForkKnowledgeFileResponseBody) GoString() string {
	return s.String()
}

func (s *ForkKnowledgeFileResponseBody) GetData() *ForkKnowledgeFileResponseBodyData {
	return s.Data
}

func (s *ForkKnowledgeFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ForkKnowledgeFileResponseBody) SetData(v *ForkKnowledgeFileResponseBodyData) *ForkKnowledgeFileResponseBody {
	s.Data = v
	return s
}

func (s *ForkKnowledgeFileResponseBody) SetRequestId(v string) *ForkKnowledgeFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForkKnowledgeFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ForkKnowledgeFileResponseBodyData struct {
	// The ID of the newly created reference record, represented as a decimal string. Use this ID for subsequent management of the record. This parameter is returned only when the operation is successful.
	//
	// example:
	//
	// 456
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The address of the new reference record, in the format of fork://<source file ID>/<new file ID>. This parameter is returned only when the operation is successful.
	//
	// example:
	//
	// fork://123/456
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The description of the operation result. If the business operation fails, this parameter indicates the cause or the security error code.
	//
	// example:
	//
	// Fork succeeded
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the content host file, represented as a decimal string. Currently, this value is the same as the source file ID. This parameter is returned only when the operation is successful.
	//
	// example:
	//
	// 123
	OwnerFileId *string `json:"OwnerFileId,omitempty" xml:"OwnerFileId,omitempty"`
	// The status of the new reference record. The value is fixed to DUPLICATED. This parameter is returned only when the operation is successful.
	//
	// example:
	//
	// DUPLICATED
	ProcessStatus *string `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// The ID of the source file, represented as a decimal string. This parameter is returned only when the operation is successful.
	//
	// example:
	//
	// 123
	SourceFileId *string `json:"SourceFileId,omitempty" xml:"SourceFileId,omitempty"`
	// Indicates whether the business operation is successful. If the operation fails, the Data parameter contains only Success and Message.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ForkKnowledgeFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ForkKnowledgeFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *ForkKnowledgeFileResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *ForkKnowledgeFileResponseBodyData) GetFileLocation() *string {
	return s.FileLocation
}

func (s *ForkKnowledgeFileResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ForkKnowledgeFileResponseBodyData) GetOwnerFileId() *string {
	return s.OwnerFileId
}

func (s *ForkKnowledgeFileResponseBodyData) GetProcessStatus() *string {
	return s.ProcessStatus
}

func (s *ForkKnowledgeFileResponseBodyData) GetSourceFileId() *string {
	return s.SourceFileId
}

func (s *ForkKnowledgeFileResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ForkKnowledgeFileResponseBodyData) SetFileId(v string) *ForkKnowledgeFileResponseBodyData {
	s.FileId = &v
	return s
}

func (s *ForkKnowledgeFileResponseBodyData) SetFileLocation(v string) *ForkKnowledgeFileResponseBodyData {
	s.FileLocation = &v
	return s
}

func (s *ForkKnowledgeFileResponseBodyData) SetMessage(v string) *ForkKnowledgeFileResponseBodyData {
	s.Message = &v
	return s
}

func (s *ForkKnowledgeFileResponseBodyData) SetOwnerFileId(v string) *ForkKnowledgeFileResponseBodyData {
	s.OwnerFileId = &v
	return s
}

func (s *ForkKnowledgeFileResponseBodyData) SetProcessStatus(v string) *ForkKnowledgeFileResponseBodyData {
	s.ProcessStatus = &v
	return s
}

func (s *ForkKnowledgeFileResponseBodyData) SetSourceFileId(v string) *ForkKnowledgeFileResponseBodyData {
	s.SourceFileId = &v
	return s
}

func (s *ForkKnowledgeFileResponseBodyData) SetSuccess(v bool) *ForkKnowledgeFileResponseBodyData {
	s.Success = &v
	return s
}

func (s *ForkKnowledgeFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
