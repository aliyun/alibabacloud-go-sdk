// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGroupDirectoryResponseBody
	GetCode() *string
	SetDescription(v string) *UpdateGroupDirectoryResponseBody
	GetDescription() *string
	SetDirectoryId(v string) *UpdateGroupDirectoryResponseBody
	GetDirectoryId() *string
	SetDirectoryType(v string) *UpdateGroupDirectoryResponseBody
	GetDirectoryType() *string
	SetGroupId(v string) *UpdateGroupDirectoryResponseBody
	GetGroupId() *string
	SetKbRootDirectoryId(v string) *UpdateGroupDirectoryResponseBody
	GetKbRootDirectoryId() *string
	SetMessage(v string) *UpdateGroupDirectoryResponseBody
	GetMessage() *string
	SetName(v string) *UpdateGroupDirectoryResponseBody
	GetName() *string
	SetParentDirectoryId(v string) *UpdateGroupDirectoryResponseBody
	GetParentDirectoryId() *string
	SetRequestId(v string) *UpdateGroupDirectoryResponseBody
	GetRequestId() *string
}

type UpdateGroupDirectoryResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The folder description.
	//
	// example:
	//
	// Project description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// dir_example
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The folder type. The value is fixed as GROUP.
	//
	// example:
	//
	// GROUP
	DirectoryType *string `json:"directoryType,omitempty" xml:"directoryType,omitempty"`
	// The ID of the collaborative share.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The ID of the internal root folder in the collaborative share. This folder cannot be modified.
	//
	// example:
	//
	// dir_root
	KbRootDirectoryId *string `json:"kbRootDirectoryId,omitempty" xml:"kbRootDirectoryId,omitempty"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The folder name.
	//
	// example:
	//
	// Project Materials
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parent folder ID.
	//
	// example:
	//
	// dir_parent
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGroupDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupDirectoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGroupDirectoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateGroupDirectoryResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateGroupDirectoryResponseBody) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *UpdateGroupDirectoryResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateGroupDirectoryResponseBody) GetKbRootDirectoryId() *string {
	return s.KbRootDirectoryId
}

func (s *UpdateGroupDirectoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGroupDirectoryResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateGroupDirectoryResponseBody) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *UpdateGroupDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGroupDirectoryResponseBody) SetCode(v string) *UpdateGroupDirectoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGroupDirectoryResponseBody) SetDescription(v string) *UpdateGroupDirectoryResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateGroupDirectoryResponseBody) SetDirectoryId(v string) *UpdateGroupDirectoryResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *UpdateGroupDirectoryResponseBody) SetDirectoryType(v string) *UpdateGroupDirectoryResponseBody {
	s.DirectoryType = &v
	return s
}

func (s *UpdateGroupDirectoryResponseBody) SetGroupId(v string) *UpdateGroupDirectoryResponseBody {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupDirectoryResponseBody) SetKbRootDirectoryId(v string) *UpdateGroupDirectoryResponseBody {
	s.KbRootDirectoryId = &v
	return s
}

func (s *UpdateGroupDirectoryResponseBody) SetMessage(v string) *UpdateGroupDirectoryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGroupDirectoryResponseBody) SetName(v string) *UpdateGroupDirectoryResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateGroupDirectoryResponseBody) SetParentDirectoryId(v string) *UpdateGroupDirectoryResponseBody {
	s.ParentDirectoryId = &v
	return s
}

func (s *UpdateGroupDirectoryResponseBody) SetRequestId(v string) *UpdateGroupDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
