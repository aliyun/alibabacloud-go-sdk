// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeUploadUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListKnowledgeUploadUserResponseBodyData) *ListKnowledgeUploadUserResponseBody
	GetData() *ListKnowledgeUploadUserResponseBodyData
	SetRequestId(v string) *ListKnowledgeUploadUserResponseBody
	GetRequestId() *string
}

type ListKnowledgeUploadUserResponseBody struct {
	// The returned data.
	Data *ListKnowledgeUploadUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListKnowledgeUploadUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeUploadUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeUploadUserResponseBody) GetData() *ListKnowledgeUploadUserResponseBodyData {
	return s.Data
}

func (s *ListKnowledgeUploadUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeUploadUserResponseBody) SetData(v *ListKnowledgeUploadUserResponseBodyData) *ListKnowledgeUploadUserResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeUploadUserResponseBody) SetRequestId(v string) *ListKnowledgeUploadUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeUploadUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeUploadUserResponseBodyData struct {
	// The location of the knowledge base file.
	//
	// example:
	//
	// oss://bucket/doc.pdf
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
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
	// The list of authorized users.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListKnowledgeUploadUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeUploadUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeUploadUserResponseBodyData) GetFileLocation() *string {
	return s.FileLocation
}

func (s *ListKnowledgeUploadUserResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeUploadUserResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeUploadUserResponseBodyData) GetUsers() []*string {
	return s.Users
}

func (s *ListKnowledgeUploadUserResponseBodyData) SetFileLocation(v string) *ListKnowledgeUploadUserResponseBodyData {
	s.FileLocation = &v
	return s
}

func (s *ListKnowledgeUploadUserResponseBodyData) SetMessage(v string) *ListKnowledgeUploadUserResponseBodyData {
	s.Message = &v
	return s
}

func (s *ListKnowledgeUploadUserResponseBodyData) SetSuccess(v bool) *ListKnowledgeUploadUserResponseBodyData {
	s.Success = &v
	return s
}

func (s *ListKnowledgeUploadUserResponseBodyData) SetUsers(v []*string) *ListKnowledgeUploadUserResponseBodyData {
	s.Users = v
	return s
}

func (s *ListKnowledgeUploadUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
