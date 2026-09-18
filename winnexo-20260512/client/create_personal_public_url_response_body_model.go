// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalPublicUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreatePersonalPublicUrlResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalPublicUrlResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalPublicUrlResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalPublicUrlResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalPublicUrlResponseBody
	GetName() *string
	SetOriginalUrl(v string) *CreatePersonalPublicUrlResponseBody
	GetOriginalUrl() *string
	SetRequestId(v string) *CreatePersonalPublicUrlResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalPublicUrlResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalPublicUrlResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalPublicUrlResponseBody
	GetStatus() *string
}

type CreatePersonalPublicUrlResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The ID of the target personal directory specified in the request. This value is empty when the default root directory is used.
	//
	// example:
	//
	// dir_personal_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2026-09-14T10:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The error description. This value is empty if the request is successful.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The resource name.
	//
	// example:
	//
	// Project Plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The submitted public web page URL.
	//
	// example:
	//
	// https://example.com
	OriginalUrl *string `json:"originalUrl,omitempty" xml:"originalUrl,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The resource scope. This value is fixed to PERSONAL.
	//
	// example:
	//
	// PERSONAL
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The ID of the newly created source.
	//
	// example:
	//
	// src_public_url_1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The resource status. A value of RUNNING indicates that the request has been accepted but crawling and parsing are not yet complete.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalPublicUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalPublicUrlResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalPublicUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalPublicUrlResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalPublicUrlResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalPublicUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalPublicUrlResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalPublicUrlResponseBody) GetOriginalUrl() *string {
	return s.OriginalUrl
}

func (s *CreatePersonalPublicUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalPublicUrlResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalPublicUrlResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalPublicUrlResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalPublicUrlResponseBody) SetCode(v string) *CreatePersonalPublicUrlResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalPublicUrlResponseBody) SetDirectoryId(v string) *CreatePersonalPublicUrlResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalPublicUrlResponseBody) SetGmtCreate(v string) *CreatePersonalPublicUrlResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalPublicUrlResponseBody) SetMessage(v string) *CreatePersonalPublicUrlResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalPublicUrlResponseBody) SetName(v string) *CreatePersonalPublicUrlResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalPublicUrlResponseBody) SetOriginalUrl(v string) *CreatePersonalPublicUrlResponseBody {
	s.OriginalUrl = &v
	return s
}

func (s *CreatePersonalPublicUrlResponseBody) SetRequestId(v string) *CreatePersonalPublicUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalPublicUrlResponseBody) SetScope(v string) *CreatePersonalPublicUrlResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalPublicUrlResponseBody) SetSourceId(v string) *CreatePersonalPublicUrlResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalPublicUrlResponseBody) SetStatus(v string) *CreatePersonalPublicUrlResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalPublicUrlResponseBody) Validate() error {
	return dara.Validate(s)
}
