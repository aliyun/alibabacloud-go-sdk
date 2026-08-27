// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalFileResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalFileResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalFileResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalFileResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalFileResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalFileResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalFileResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalFileResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalFileResponseBody
	GetStatus() *string
}

type CreatePersonalFileResponseBody struct {
	// The response status code.
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
	// The error details.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The updated filter view name.
	//
	// example:
	//
	// ha-cn-36z45q4xg06_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3DAC4165-2401-543B-B5E7-A86AA151E517
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The export scope. Valid values:
	//
	// - ALL: all.
	//
	// - SELECT: selected rows.
	//
	// example:
	//
	// read:user,read:repo,write:repo,read:org,read:group
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The source ID.
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

func (s CreatePersonalFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalFileResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFileResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalFileResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalFileResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalFileResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalFileResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalFileResponseBody) SetCode(v string) *CreatePersonalFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetDirectoryId(v string) *CreatePersonalFileResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetGmtCreate(v string) *CreatePersonalFileResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetMessage(v string) *CreatePersonalFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetName(v string) *CreatePersonalFileResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetRequestId(v string) *CreatePersonalFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetScope(v string) *CreatePersonalFileResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetSourceId(v string) *CreatePersonalFileResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalFileResponseBody) SetStatus(v string) *CreatePersonalFileResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalFileResponseBody) Validate() error {
	return dara.Validate(s)
}
