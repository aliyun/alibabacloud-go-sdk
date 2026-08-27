// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MoveResourceResponseBody
	GetCode() *string
	SetMessage(v string) *MoveResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *MoveResourceResponseBody
	GetRequestId() *string
	SetSourceDirectoryId(v string) *MoveResourceResponseBody
	GetSourceDirectoryId() *string
	SetSourceId(v string) *MoveResourceResponseBody
	GetSourceId() *string
	SetSuccess(v bool) *MoveResourceResponseBody
	GetSuccess() *bool
	SetTargetDirectoryId(v string) *MoveResourceResponseBody
	GetTargetDirectoryId() *string
}

type MoveResourceResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E68654BD-F7BA-5837-8686-5645D739A47C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The source directory ID, which echoes the input parameter.
	//
	// example:
	//
	// exampleSourceDirectoryId
	SourceDirectoryId *string `json:"sourceDirectoryId,omitempty" xml:"sourceDirectoryId,omitempty"`
	// The resource ID, which echoes the input parameter.
	//
	// example:
	//
	// 2000627
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The target directory ID.
	//
	// example:
	//
	// exampleTargetDirectoryId
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
}

func (s MoveResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MoveResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *MoveResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MoveResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MoveResourceResponseBody) GetSourceDirectoryId() *string {
	return s.SourceDirectoryId
}

func (s *MoveResourceResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *MoveResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MoveResourceResponseBody) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *MoveResourceResponseBody) SetCode(v string) *MoveResourceResponseBody {
	s.Code = &v
	return s
}

func (s *MoveResourceResponseBody) SetMessage(v string) *MoveResourceResponseBody {
	s.Message = &v
	return s
}

func (s *MoveResourceResponseBody) SetRequestId(v string) *MoveResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceResponseBody) SetSourceDirectoryId(v string) *MoveResourceResponseBody {
	s.SourceDirectoryId = &v
	return s
}

func (s *MoveResourceResponseBody) SetSourceId(v string) *MoveResourceResponseBody {
	s.SourceId = &v
	return s
}

func (s *MoveResourceResponseBody) SetSuccess(v bool) *MoveResourceResponseBody {
	s.Success = &v
	return s
}

func (s *MoveResourceResponseBody) SetTargetDirectoryId(v string) *MoveResourceResponseBody {
	s.TargetDirectoryId = &v
	return s
}

func (s *MoveResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
