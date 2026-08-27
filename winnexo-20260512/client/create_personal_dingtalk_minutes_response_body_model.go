// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkMinutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalDingtalkMinutesResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalDingtalkMinutesResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalDingtalkMinutesResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalDingtalkMinutesResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalDingtalkMinutesResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalDingtalkMinutesResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalDingtalkMinutesResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalDingtalkMinutesResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalDingtalkMinutesResponseBody
	GetStatus() *string
}

type CreatePersonalDingtalkMinutesResponseBody struct {
	// The error code.
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
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the AI assistant.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 72D1EC35-B174-5595-891F-2F0B3BFBE56F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// read:user,read:repo,write:repo,read:org,read:group
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The unique identifier on the business system side, which is the business ID.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The status.
	//
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalDingtalkMinutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkMinutesResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkMinutesResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalDingtalkMinutesResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalDingtalkMinutesResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalDingtalkMinutesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalDingtalkMinutesResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalDingtalkMinutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalDingtalkMinutesResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalDingtalkMinutesResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalDingtalkMinutesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalDingtalkMinutesResponseBody) SetCode(v string) *CreatePersonalDingtalkMinutesResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponseBody) SetDirectoryId(v string) *CreatePersonalDingtalkMinutesResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponseBody) SetGmtCreate(v string) *CreatePersonalDingtalkMinutesResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponseBody) SetMessage(v string) *CreatePersonalDingtalkMinutesResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponseBody) SetName(v string) *CreatePersonalDingtalkMinutesResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponseBody) SetRequestId(v string) *CreatePersonalDingtalkMinutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponseBody) SetScope(v string) *CreatePersonalDingtalkMinutesResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponseBody) SetSourceId(v string) *CreatePersonalDingtalkMinutesResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponseBody) SetStatus(v string) *CreatePersonalDingtalkMinutesResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesResponseBody) Validate() error {
	return dara.Validate(s)
}
