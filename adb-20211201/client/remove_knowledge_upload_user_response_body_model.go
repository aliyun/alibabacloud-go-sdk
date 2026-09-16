// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveKnowledgeUploadUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RemoveKnowledgeUploadUserResponseBodyData) *RemoveKnowledgeUploadUserResponseBody
	GetData() *RemoveKnowledgeUploadUserResponseBodyData
	SetRequestId(v string) *RemoveKnowledgeUploadUserResponseBody
	GetRequestId() *string
}

type RemoveKnowledgeUploadUserResponseBody struct {
	// The returned data.
	Data *RemoveKnowledgeUploadUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveKnowledgeUploadUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeUploadUserResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeUploadUserResponseBody) GetData() *RemoveKnowledgeUploadUserResponseBodyData {
	return s.Data
}

func (s *RemoveKnowledgeUploadUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveKnowledgeUploadUserResponseBody) SetData(v *RemoveKnowledgeUploadUserResponseBodyData) *RemoveKnowledgeUploadUserResponseBody {
	s.Data = v
	return s
}

func (s *RemoveKnowledgeUploadUserResponseBody) SetRequestId(v string) *RemoveKnowledgeUploadUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveKnowledgeUploadUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveKnowledgeUploadUserResponseBodyData struct {
	// The location of the knowledge base file.
	//
	// example:
	//
	// oss://bucketName/path/to/file.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of users that were successfully deleted.
	//
	// example:
	//
	// 1
	Removed *int32 `json:"Removed,omitempty" xml:"Removed,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveKnowledgeUploadUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeUploadUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeUploadUserResponseBodyData) GetFileLocation() *string {
	return s.FileLocation
}

func (s *RemoveKnowledgeUploadUserResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *RemoveKnowledgeUploadUserResponseBodyData) GetRemoved() *int32 {
	return s.Removed
}

func (s *RemoveKnowledgeUploadUserResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveKnowledgeUploadUserResponseBodyData) SetFileLocation(v string) *RemoveKnowledgeUploadUserResponseBodyData {
	s.FileLocation = &v
	return s
}

func (s *RemoveKnowledgeUploadUserResponseBodyData) SetMessage(v string) *RemoveKnowledgeUploadUserResponseBodyData {
	s.Message = &v
	return s
}

func (s *RemoveKnowledgeUploadUserResponseBodyData) SetRemoved(v int32) *RemoveKnowledgeUploadUserResponseBodyData {
	s.Removed = &v
	return s
}

func (s *RemoveKnowledgeUploadUserResponseBodyData) SetSuccess(v bool) *RemoveKnowledgeUploadUserResponseBodyData {
	s.Success = &v
	return s
}

func (s *RemoveKnowledgeUploadUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
