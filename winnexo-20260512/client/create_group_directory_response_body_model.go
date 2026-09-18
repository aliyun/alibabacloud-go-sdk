// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGroupDirectoryResponseBody
	GetCode() *string
	SetDescription(v string) *CreateGroupDirectoryResponseBody
	GetDescription() *string
	SetDirectoryId(v string) *CreateGroupDirectoryResponseBody
	GetDirectoryId() *string
	SetDirectoryType(v string) *CreateGroupDirectoryResponseBody
	GetDirectoryType() *string
	SetGroupId(v string) *CreateGroupDirectoryResponseBody
	GetGroupId() *string
	SetKbRootDirectoryId(v string) *CreateGroupDirectoryResponseBody
	GetKbRootDirectoryId() *string
	SetMessage(v string) *CreateGroupDirectoryResponseBody
	GetMessage() *string
	SetName(v string) *CreateGroupDirectoryResponseBody
	GetName() *string
	SetParentDirectoryId(v string) *CreateGroupDirectoryResponseBody
	GetParentDirectoryId() *string
	SetRequestId(v string) *CreateGroupDirectoryResponseBody
	GetRequestId() *string
}

type CreateGroupDirectoryResponseBody struct {
	// SUCCESS indicates success. In failure cases, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the AI assistant.
	//
	// example:
	//
	// ProjectDescription
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The folder type.
	//
	// example:
	//
	// GROUP
	DirectoryType *string `json:"directoryType,omitempty" xml:"directoryType,omitempty"`
	// The project group ID.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The root folder ID of the knowledge base.
	//
	// example:
	//
	// dir_root
	KbRootDirectoryId *string `json:"kbRootDirectoryId,omitempty" xml:"kbRootDirectoryId,omitempty"`
	// The response message.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the worksheet.
	//
	// example:
	//
	// ProjectFiles
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// dir_parent
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F4A9EB1C-6952-5CCC-B1DC-355576FC82A7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateGroupDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGroupDirectoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateGroupDirectoryResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupDirectoryResponseBody) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *CreateGroupDirectoryResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupDirectoryResponseBody) GetKbRootDirectoryId() *string {
	return s.KbRootDirectoryId
}

func (s *CreateGroupDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupDirectoryResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateGroupDirectoryResponseBody) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreateGroupDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupDirectoryResponseBody) SetCode(v string) *CreateGroupDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupDirectoryResponseBody) SetDescription(v string) *CreateGroupDirectoryResponseBody {
	s.Description = &v
	return s
}

func (s *CreateGroupDirectoryResponseBody) SetDirectoryId(v string) *CreateGroupDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupDirectoryResponseBody) SetDirectoryType(v string) *CreateGroupDirectoryResponseBody {
	s.DirectoryType = &v
	return s
}

func (s *CreateGroupDirectoryResponseBody) SetGroupId(v string) *CreateGroupDirectoryResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupDirectoryResponseBody) SetKbRootDirectoryId(v string) *CreateGroupDirectoryResponseBody {
	s.KbRootDirectoryId = &v
	return s
}

func (s *CreateGroupDirectoryResponseBody) SetMessage(v string) *CreateGroupDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupDirectoryResponseBody) SetName(v string) *CreateGroupDirectoryResponseBody {
	s.Name = &v
	return s
}

func (s *CreateGroupDirectoryResponseBody) SetParentDirectoryId(v string) *CreateGroupDirectoryResponseBody {
	s.ParentDirectoryId = &v
	return s
}

func (s *CreateGroupDirectoryResponseBody) SetRequestId(v string) *CreateGroupDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
