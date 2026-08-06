// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKgSchemaPublishResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKgSchemaPublishResultResponseBody
	GetCode() *string
	SetData(v *GetKgSchemaPublishResultResponseBodyData) *GetKgSchemaPublishResultResponseBody
	GetData() *GetKgSchemaPublishResultResponseBodyData
	SetHttpStatusCode(v int32) *GetKgSchemaPublishResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetKgSchemaPublishResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKgSchemaPublishResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKgSchemaPublishResultResponseBody
	GetSuccess() *bool
}

type GetKgSchemaPublishResultResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The publish result.
	Data *GetKgSchemaPublishResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetKgSchemaPublishResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKgSchemaPublishResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetKgSchemaPublishResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKgSchemaPublishResultResponseBody) GetData() *GetKgSchemaPublishResultResponseBodyData {
	return s.Data
}

func (s *GetKgSchemaPublishResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetKgSchemaPublishResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKgSchemaPublishResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKgSchemaPublishResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKgSchemaPublishResultResponseBody) SetCode(v string) *GetKgSchemaPublishResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetKgSchemaPublishResultResponseBody) SetData(v *GetKgSchemaPublishResultResponseBodyData) *GetKgSchemaPublishResultResponseBody {
	s.Data = v
	return s
}

func (s *GetKgSchemaPublishResultResponseBody) SetHttpStatusCode(v int32) *GetKgSchemaPublishResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetKgSchemaPublishResultResponseBody) SetMessage(v string) *GetKgSchemaPublishResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetKgSchemaPublishResultResponseBody) SetRequestId(v string) *GetKgSchemaPublishResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKgSchemaPublishResultResponseBody) SetSuccess(v bool) *GetKgSchemaPublishResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetKgSchemaPublishResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKgSchemaPublishResultResponseBodyData struct {
	// The publish content.
	//
	// example:
	//
	// test
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The publish status. Valid values:
	//
	// - Publishing: Publishing in progress.
	//
	// - Published: Published successfully.
	//
	// - Partial: Partially completed.
	//
	// - Failed: Failed.
	//
	// - RollbackFailed: Publish failed and rollback failed.
	//
	// example:
	//
	// Published
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	VersionId *int32 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetKgSchemaPublishResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKgSchemaPublishResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKgSchemaPublishResultResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetKgSchemaPublishResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetKgSchemaPublishResultResponseBodyData) GetVersionId() *int32 {
	return s.VersionId
}

func (s *GetKgSchemaPublishResultResponseBodyData) SetContent(v string) *GetKgSchemaPublishResultResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetKgSchemaPublishResultResponseBodyData) SetStatus(v string) *GetKgSchemaPublishResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetKgSchemaPublishResultResponseBodyData) SetVersionId(v int32) *GetKgSchemaPublishResultResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *GetKgSchemaPublishResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
