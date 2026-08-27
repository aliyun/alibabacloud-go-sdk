// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuMinuteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalFeishuMinuteResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalFeishuMinuteResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalFeishuMinuteResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalFeishuMinuteResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalFeishuMinuteResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalFeishuMinuteResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalFeishuMinuteResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalFeishuMinuteResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalFeishuMinuteResponseBody
	GetStatus() *string
}

type CreatePersonalFeishuMinuteResponseBody struct {
	// SUCCESS indicates success. In case of failure, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-03-04 11:12:03
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The operation message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The resource name.
	//
	// example:
	//
	// issue_research
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FEE93-17FB-5369-BB65-1188C3A14B0A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// read:user,read:repo,write:repo,read:org,read:group
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The original project ID.
	//
	// example:
	//
	// 2001086
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The status.
	//
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalFeishuMinuteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuMinuteResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalFeishuMinuteResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetCode(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetDirectoryId(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetGmtCreate(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetMessage(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetName(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetRequestId(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetScope(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetSourceId(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) SetStatus(v string) *CreatePersonalFeishuMinuteResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalFeishuMinuteResponseBody) Validate() error {
	return dara.Validate(s)
}
