// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkMeetingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalDingtalkMeetingResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalDingtalkMeetingResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalDingtalkMeetingResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalDingtalkMeetingResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalDingtalkMeetingResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalDingtalkMeetingResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalDingtalkMeetingResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalDingtalkMeetingResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalDingtalkMeetingResponseBody
	GetStatus() *string
}

type CreatePersonalDingtalkMeetingResponseBody struct {
	// The status code.
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
	// 2025-11-14T02:18:27Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The response message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The skill name.
	//
	// example:
	//
	// oklabs_tongyici
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B49109FE-5BB1-593C-915D-F5A99D9F5435
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// user_info projects pull_requests hook gists emails
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The unique identifier on the business system side, which is the business ID.
	//
	// example:
	//
	// 2000398
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The task running status.
	//
	// example:
	//
	// updated
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalDingtalkMeetingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkMeetingResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalDingtalkMeetingResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalDingtalkMeetingResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalDingtalkMeetingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalDingtalkMeetingResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalDingtalkMeetingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalDingtalkMeetingResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalDingtalkMeetingResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalDingtalkMeetingResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalDingtalkMeetingResponseBody) SetCode(v string) *CreatePersonalDingtalkMeetingResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponseBody) SetDirectoryId(v string) *CreatePersonalDingtalkMeetingResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponseBody) SetGmtCreate(v string) *CreatePersonalDingtalkMeetingResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponseBody) SetMessage(v string) *CreatePersonalDingtalkMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponseBody) SetName(v string) *CreatePersonalDingtalkMeetingResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponseBody) SetRequestId(v string) *CreatePersonalDingtalkMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponseBody) SetScope(v string) *CreatePersonalDingtalkMeetingResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponseBody) SetSourceId(v string) *CreatePersonalDingtalkMeetingResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponseBody) SetStatus(v string) *CreatePersonalDingtalkMeetingResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalDingtalkMeetingResponseBody) Validate() error {
	return dara.Validate(s)
}
