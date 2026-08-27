// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalVoiceMeetingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalVoiceMeetingResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalVoiceMeetingResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalVoiceMeetingResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalVoiceMeetingResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalVoiceMeetingResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalVoiceMeetingResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalVoiceMeetingResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalVoiceMeetingResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalVoiceMeetingResponseBody
	GetStatus() *string
}

type CreatePersonalVoiceMeetingResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The bound directory ID. This value echoes the directoryId provided in the request body. If no directoryId is specified, the value is null because the default root directory is used.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2026-04-22T08:15:28.000+00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The knowledge base name.
	//
	// example:
	//
	// p-toolset-e95d1287-3d40-487a-bcce-6e6252c7a793
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5241B90-8FF4-565C-977A-0CE1842AED72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The resource scope. The value is fixed to PERSONAL.
	//
	// example:
	//
	// user_info projects pull_requests hook gists emails
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The unique identifier on the business system side, that is, the business ID.
	//
	// example:
	//
	// 8
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The resource status. The initial status during the creation process is typically PENDING. If the on_create process fails, the status is FAILED.
	//
	// example:
	//
	// {\\"observedGeneration\\": 4, \\"servicesInstances\\": {}, \\"observedTime\\": \\"2025-10-31T03:48:27Z\\", \\"servicesWithPendingChanges\\": [], \\"latestEnvironmentDeploymentName\\": \\"manual-1761882507097-Eu1vIP\\"}
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalVoiceMeetingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalVoiceMeetingResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalVoiceMeetingResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetCode(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetDirectoryId(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetGmtCreate(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetMessage(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetName(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetRequestId(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetScope(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetSourceId(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) SetStatus(v string) *CreatePersonalVoiceMeetingResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalVoiceMeetingResponseBody) Validate() error {
	return dara.Validate(s)
}
