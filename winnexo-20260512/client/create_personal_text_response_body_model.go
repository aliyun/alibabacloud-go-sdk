// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalTextResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalTextResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalTextResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalTextResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalTextResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalTextResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalTextResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalTextResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalTextResponseBody
	GetStatus() *string
}

type CreatePersonalTextResponseBody struct {
	// SUCCESS indicates success. In case of failure, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2026-05-22 18:18:56
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The response message.
	//
	// example:
	//
	// Instance i-0jl6hlcbtuo4eqg7puni not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The updated filter view name.
	//
	// example:
	//
	// p-toolset-3dcef7ca-31b9-4d1c-8692-1ef03099cad3
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04EE99E6-A0D9-5B04-81D1-7BEC0CB0AFDF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The travel scale of the integration partner.
	//
	// example:
	//
	// read:user,read:repo,write:repo,read:org,read:group
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The source ID.
	//
	// example:
	//
	// 2000398
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The filter status.
	//
	// - 0: All
	//
	// - 1: Unconfirmed
	//
	// - 3: Ignored
	//
	// - 4: Rejected
	//
	// example:
	//
	// {\\"observedGeneration\\": 4, \\"servicesInstances\\": {}, \\"observedTime\\": \\"2025-10-31T03:48:27Z\\", \\"servicesWithPendingChanges\\": [], \\"latestEnvironmentDeploymentName\\": \\"manual-1761882507097-Eu1vIP\\"}
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalTextResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalTextResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalTextResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalTextResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalTextResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalTextResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalTextResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalTextResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalTextResponseBody) SetCode(v string) *CreatePersonalTextResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetDirectoryId(v string) *CreatePersonalTextResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetGmtCreate(v string) *CreatePersonalTextResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetMessage(v string) *CreatePersonalTextResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetName(v string) *CreatePersonalTextResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetRequestId(v string) *CreatePersonalTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetScope(v string) *CreatePersonalTextResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetSourceId(v string) *CreatePersonalTextResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalTextResponseBody) SetStatus(v string) *CreatePersonalTextResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalTextResponseBody) Validate() error {
	return dara.Validate(s)
}
