// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishKgSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PublishKgSchemaResponseBody
	GetCode() *string
	SetData(v *PublishKgSchemaResponseBodyData) *PublishKgSchemaResponseBody
	GetData() *PublishKgSchemaResponseBodyData
	SetHttpStatusCode(v int32) *PublishKgSchemaResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PublishKgSchemaResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishKgSchemaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PublishKgSchemaResponseBody
	GetSuccess() *bool
}

type PublishKgSchemaResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The publish result.
	Data *PublishKgSchemaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PublishKgSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishKgSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *PublishKgSchemaResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishKgSchemaResponseBody) GetData() *PublishKgSchemaResponseBodyData {
	return s.Data
}

func (s *PublishKgSchemaResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PublishKgSchemaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishKgSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishKgSchemaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PublishKgSchemaResponseBody) SetCode(v string) *PublishKgSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *PublishKgSchemaResponseBody) SetData(v *PublishKgSchemaResponseBodyData) *PublishKgSchemaResponseBody {
	s.Data = v
	return s
}

func (s *PublishKgSchemaResponseBody) SetHttpStatusCode(v int32) *PublishKgSchemaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PublishKgSchemaResponseBody) SetMessage(v string) *PublishKgSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *PublishKgSchemaResponseBody) SetRequestId(v string) *PublishKgSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishKgSchemaResponseBody) SetSuccess(v bool) *PublishKgSchemaResponseBody {
	s.Success = &v
	return s
}

func (s *PublishKgSchemaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishKgSchemaResponseBodyData struct {
	// The expected latest model version number after the publish operation.
	//
	// example:
	//
	// 1
	VersionId *int32 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The model ID.
	//
	// example:
	//
	// f1d4559a4db044158305e2d89bccf81f
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PublishKgSchemaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PublishKgSchemaResponseBodyData) GoString() string {
	return s.String()
}

func (s *PublishKgSchemaResponseBodyData) GetVersionId() *int32 {
	return s.VersionId
}

func (s *PublishKgSchemaResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PublishKgSchemaResponseBodyData) SetVersionId(v int32) *PublishKgSchemaResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *PublishKgSchemaResponseBodyData) SetWorkspaceId(v string) *PublishKgSchemaResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *PublishKgSchemaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
