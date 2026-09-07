// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeUploadUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddKnowledgeUploadUserResponseBodyData) *AddKnowledgeUploadUserResponseBody
	GetData() *AddKnowledgeUploadUserResponseBodyData
	SetRequestId(v string) *AddKnowledgeUploadUserResponseBody
	GetRequestId() *string
}

type AddKnowledgeUploadUserResponseBody struct {
	// The returned data.
	Data *AddKnowledgeUploadUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddKnowledgeUploadUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeUploadUserResponseBody) GoString() string {
	return s.String()
}

func (s *AddKnowledgeUploadUserResponseBody) GetData() *AddKnowledgeUploadUserResponseBodyData {
	return s.Data
}

func (s *AddKnowledgeUploadUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddKnowledgeUploadUserResponseBody) SetData(v *AddKnowledgeUploadUserResponseBodyData) *AddKnowledgeUploadUserResponseBody {
	s.Data = v
	return s
}

func (s *AddKnowledgeUploadUserResponseBody) SetRequestId(v string) *AddKnowledgeUploadUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddKnowledgeUploadUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddKnowledgeUploadUserResponseBodyData struct {
	// The location of the knowledge base file.
	//
	// example:
	//
	// oss://bucketName/path/to/file.pfg
	FileLocation *string `json:"FileLocation,omitempty" xml:"FileLocation,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of skipped users.
	Skipped []*AddKnowledgeUploadUserResponseBodyDataSkipped `json:"Skipped,omitempty" xml:"Skipped,omitempty" type:"Repeated"`
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
	// The number of users that were successfully added.
	//
	// example:
	//
	// 1
	Written *int32 `json:"Written,omitempty" xml:"Written,omitempty"`
}

func (s AddKnowledgeUploadUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeUploadUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddKnowledgeUploadUserResponseBodyData) GetFileLocation() *string {
	return s.FileLocation
}

func (s *AddKnowledgeUploadUserResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *AddKnowledgeUploadUserResponseBodyData) GetSkipped() []*AddKnowledgeUploadUserResponseBodyDataSkipped {
	return s.Skipped
}

func (s *AddKnowledgeUploadUserResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *AddKnowledgeUploadUserResponseBodyData) GetWritten() *int32 {
	return s.Written
}

func (s *AddKnowledgeUploadUserResponseBodyData) SetFileLocation(v string) *AddKnowledgeUploadUserResponseBodyData {
	s.FileLocation = &v
	return s
}

func (s *AddKnowledgeUploadUserResponseBodyData) SetMessage(v string) *AddKnowledgeUploadUserResponseBodyData {
	s.Message = &v
	return s
}

func (s *AddKnowledgeUploadUserResponseBodyData) SetSkipped(v []*AddKnowledgeUploadUserResponseBodyDataSkipped) *AddKnowledgeUploadUserResponseBodyData {
	s.Skipped = v
	return s
}

func (s *AddKnowledgeUploadUserResponseBodyData) SetSuccess(v bool) *AddKnowledgeUploadUserResponseBodyData {
	s.Success = &v
	return s
}

func (s *AddKnowledgeUploadUserResponseBodyData) SetWritten(v int32) *AddKnowledgeUploadUserResponseBodyData {
	s.Written = &v
	return s
}

func (s *AddKnowledgeUploadUserResponseBodyData) Validate() error {
	if s.Skipped != nil {
		for _, item := range s.Skipped {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddKnowledgeUploadUserResponseBodyDataSkipped struct {
	// The reason why the user was skipped.
	//
	// example:
	//
	// conflicts error
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The username of the authorized user.
	//
	// example:
	//
	// test_user
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s AddKnowledgeUploadUserResponseBodyDataSkipped) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeUploadUserResponseBodyDataSkipped) GoString() string {
	return s.String()
}

func (s *AddKnowledgeUploadUserResponseBodyDataSkipped) GetReason() *string {
	return s.Reason
}

func (s *AddKnowledgeUploadUserResponseBodyDataSkipped) GetUser() *string {
	return s.User
}

func (s *AddKnowledgeUploadUserResponseBodyDataSkipped) SetReason(v string) *AddKnowledgeUploadUserResponseBodyDataSkipped {
	s.Reason = &v
	return s
}

func (s *AddKnowledgeUploadUserResponseBodyDataSkipped) SetUser(v string) *AddKnowledgeUploadUserResponseBodyDataSkipped {
	s.User = &v
	return s
}

func (s *AddKnowledgeUploadUserResponseBodyDataSkipped) Validate() error {
	return dara.Validate(s)
}
