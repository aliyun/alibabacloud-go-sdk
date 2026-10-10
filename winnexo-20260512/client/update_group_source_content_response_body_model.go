// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGroupSourceContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGroupSourceContentResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateGroupSourceContentResponseBody
	GetMessage() *string
	SetName(v string) *UpdateGroupSourceContentResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateGroupSourceContentResponseBody
	GetRequestId() *string
	SetSourceId(v string) *UpdateGroupSourceContentResponseBody
	GetSourceId() *string
	SetSourceType(v string) *UpdateGroupSourceContentResponseBody
	GetSourceType() *string
	SetStatus(v string) *UpdateGroupSourceContentResponseBody
	GetStatus() *string
}

type UpdateGroupSourceContentResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The image name.
	//
	// example:
	//
	// Project resource
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// C474BFC7-7B11-5D92-971E-74AA82EC495B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// source_example
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The data source type.
	//
	// example:
	//
	// example
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The task running status.
	//
	// example:
	//
	// example
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateGroupSourceContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGroupSourceContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGroupSourceContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGroupSourceContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGroupSourceContentResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateGroupSourceContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGroupSourceContentResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateGroupSourceContentResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateGroupSourceContentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateGroupSourceContentResponseBody) SetCode(v string) *UpdateGroupSourceContentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetMessage(v string) *UpdateGroupSourceContentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetName(v string) *UpdateGroupSourceContentResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetRequestId(v string) *UpdateGroupSourceContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetSourceId(v string) *UpdateGroupSourceContentResponseBody {
	s.SourceId = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetSourceType(v string) *UpdateGroupSourceContentResponseBody {
	s.SourceType = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) SetStatus(v string) *UpdateGroupSourceContentResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateGroupSourceContentResponseBody) Validate() error {
	return dara.Validate(s)
}
