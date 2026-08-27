// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAlidingDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalAlidingDocResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalAlidingDocResponseBody
	GetDirectoryId() *string
	SetFilePublicUrl(v string) *CreatePersonalAlidingDocResponseBody
	GetFilePublicUrl() *string
	SetGmtCreate(v string) *CreatePersonalAlidingDocResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalAlidingDocResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalAlidingDocResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalAlidingDocResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalAlidingDocResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalAlidingDocResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalAlidingDocResponseBody
	GetStatus() *string
}

type CreatePersonalAlidingDocResponseBody struct {
	// The response code.
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
	// The publicly accessible URL of the AliDing online document.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The timestamp when the customer group was created. Unit: milliseconds.
	//
	// example:
	//
	// 2026-04-22T08:15:28.000+00:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pipeline name.
	//
	// example:
	//
	// user_paswd_104
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E68654BD-F7BA-5837-8686-5645D739A47C
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
	// 2000413
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The call status. Valid values:
	//
	// - **PENDING**: Waiting for receipt.
	//
	// - **SUCCESS**: Succeeded.
	//
	// - **FAILED**: Failed.
	//
	// - **TIMEOUT**: Timed out.
	//
	// example:
	//
	// {\\"servicesInstances\\": {}, \\"servicesWithPendingChanges\\": []}
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalAlidingDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAlidingDocResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalAlidingDocResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalAlidingDocResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAlidingDocResponseBody) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *CreatePersonalAlidingDocResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalAlidingDocResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalAlidingDocResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalAlidingDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalAlidingDocResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalAlidingDocResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalAlidingDocResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalAlidingDocResponseBody) SetCode(v string) *CreatePersonalAlidingDocResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetDirectoryId(v string) *CreatePersonalAlidingDocResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetFilePublicUrl(v string) *CreatePersonalAlidingDocResponseBody {
	s.FilePublicUrl = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetGmtCreate(v string) *CreatePersonalAlidingDocResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetMessage(v string) *CreatePersonalAlidingDocResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetName(v string) *CreatePersonalAlidingDocResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetRequestId(v string) *CreatePersonalAlidingDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetScope(v string) *CreatePersonalAlidingDocResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetSourceId(v string) *CreatePersonalAlidingDocResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) SetStatus(v string) *CreatePersonalAlidingDocResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalAlidingDocResponseBody) Validate() error {
	return dara.Validate(s)
}
