// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveGroupResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MoveGroupResourceResponseBody
	GetCode() *string
	SetGroupId(v string) *MoveGroupResourceResponseBody
	GetGroupId() *string
	SetMessage(v string) *MoveGroupResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *MoveGroupResourceResponseBody
	GetRequestId() *string
	SetSourceDirectoryId(v string) *MoveGroupResourceResponseBody
	GetSourceDirectoryId() *string
	SetSourceId(v string) *MoveGroupResourceResponseBody
	GetSourceId() *string
	SetTargetDirectoryId(v string) *MoveGroupResourceResponseBody
	GetTargetDirectoryId() *string
}

type MoveGroupResourceResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The collaboration space ID.
	//
	// example:
	//
	// group_example
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The directory ID before the move.
	//
	// example:
	//
	// example
	SourceDirectoryId *string `json:"sourceDirectoryId,omitempty" xml:"sourceDirectoryId,omitempty"`
	// The ID of the moved resource. This value remains unchanged before and after the move.
	//
	// example:
	//
	// example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The directory ID after the move.
	//
	// example:
	//
	// example
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
}

func (s MoveGroupResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveGroupResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveGroupResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *MoveGroupResourceResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *MoveGroupResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MoveGroupResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveGroupResourceResponseBody) GetSourceDirectoryId() *string {
	return s.SourceDirectoryId
}

func (s *MoveGroupResourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *MoveGroupResourceResponseBody) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *MoveGroupResourceResponseBody) SetCode(v string) *MoveGroupResourceResponseBody {
	s.Code = &v
	return s
}

func (s *MoveGroupResourceResponseBody) SetGroupId(v string) *MoveGroupResourceResponseBody {
	s.GroupId = &v
	return s
}

func (s *MoveGroupResourceResponseBody) SetMessage(v string) *MoveGroupResourceResponseBody {
	s.Message = &v
	return s
}

func (s *MoveGroupResourceResponseBody) SetRequestId(v string) *MoveGroupResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveGroupResourceResponseBody) SetSourceDirectoryId(v string) *MoveGroupResourceResponseBody {
	s.SourceDirectoryId = &v
	return s
}

func (s *MoveGroupResourceResponseBody) SetSourceId(v string) *MoveGroupResourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *MoveGroupResourceResponseBody) SetTargetDirectoryId(v string) *MoveGroupResourceResponseBody {
	s.TargetDirectoryId = &v
	return s
}

func (s *MoveGroupResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
