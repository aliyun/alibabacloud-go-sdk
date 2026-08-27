// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSourceContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateSourceContentResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateSourceContentResponseBody
	GetMessage() *string
	SetName(v string) *UpdateSourceContentResponseBody
	GetName() *string
	SetRequestId(v string) *UpdateSourceContentResponseBody
	GetRequestId() *string
	SetSourceId(v string) *UpdateSourceContentResponseBody
	GetSourceId() *string
	SetSourceType(v string) *UpdateSourceContentResponseBody
	GetSourceType() *string
	SetStatus(v string) *UpdateSourceContentResponseBody
	GetStatus() *string
}

type UpdateSourceContentResponseBody struct {
	// The error code.
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
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The source type of the dictionary file. Valid values:
	//
	// - OSS: Object Storage Service (OSS).
	//
	// - ORIGIN: retains the previously uploaded dictionary.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The status.
	//
	// This parameter is required.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateSourceContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSourceContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSourceContentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateSourceContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSourceContentResponseBody) GetName() *string {
	return s.Name
}

func (s *UpdateSourceContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSourceContentResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateSourceContentResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateSourceContentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateSourceContentResponseBody) SetCode(v string) *UpdateSourceContentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSourceContentResponseBody) SetMessage(v string) *UpdateSourceContentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSourceContentResponseBody) SetName(v string) *UpdateSourceContentResponseBody {
	s.Name = &v
	return s
}

func (s *UpdateSourceContentResponseBody) SetRequestId(v string) *UpdateSourceContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSourceContentResponseBody) SetSourceId(v string) *UpdateSourceContentResponseBody {
	s.SourceId = &v
	return s
}

func (s *UpdateSourceContentResponseBody) SetSourceType(v string) *UpdateSourceContentResponseBody {
	s.SourceType = &v
	return s
}

func (s *UpdateSourceContentResponseBody) SetStatus(v string) *UpdateSourceContentResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateSourceContentResponseBody) Validate() error {
	return dara.Validate(s)
}
