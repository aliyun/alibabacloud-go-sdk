// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAliDingMeetingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalAliDingMeetingResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalAliDingMeetingResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalAliDingMeetingResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalAliDingMeetingResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalAliDingMeetingResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalAliDingMeetingResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalAliDingMeetingResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalAliDingMeetingResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalAliDingMeetingResponseBody
	GetStatus() *string
}

type CreatePersonalAliDingMeetingResponseBody struct {
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
	// 2025-11-12T03:08:56Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the bot.
	//
	// example:
	//
	// p-toolset-80a4520e-b35c-4e8b-acf7-3a01c7307522
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 247FBC97-433C-544A-BB29-98F572C06E9F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// user_info projects pull_requests hook gists emails
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The unique identifier on the business system side, that is, the business ID.
	//
	// example:
	//
	// 2000358
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The signing status. Valid values:
	//
	// - CREATED: Created but not signed.
	//
	// - SUCCESS: Signed.
	//
	// - STOP: Terminated.
	//
	// example:
	//
	// {\\"observedGeneration\\": 2, \\"servicesInstances\\": {}, \\"observedTime\\": \\"2026-03-05T16:00:09Z\\", \\"servicesWithPendingChanges\\": [], \\"latestEnvironmentDeploymentName\\": \\"manual-1772726409137-lmvsqr\\"}
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalAliDingMeetingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAliDingMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalAliDingMeetingResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetCode(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetDirectoryId(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetGmtCreate(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetMessage(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetName(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetRequestId(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetScope(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetSourceId(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) SetStatus(v string) *CreatePersonalAliDingMeetingResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalAliDingMeetingResponseBody) Validate() error {
	return dara.Validate(s)
}
