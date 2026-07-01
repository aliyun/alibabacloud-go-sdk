// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDatasetResponseBody
	GetCode() *string
	SetDatasetId(v string) *CreateDatasetResponseBody
	GetDatasetId() *string
	SetHttpStatusCode(v int32) *CreateDatasetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateDatasetResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDatasetResponseBody
	GetSuccess() *bool
}

type CreateDatasetResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The dataset ID (business primary key).
	//
	// example:
	//
	// 7280832407583104
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
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
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDatasetResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateDatasetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateDatasetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDatasetResponseBody) SetCode(v string) *CreateDatasetResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDatasetResponseBody) SetDatasetId(v string) *CreateDatasetResponseBody {
	s.DatasetId = &v
	return s
}

func (s *CreateDatasetResponseBody) SetHttpStatusCode(v int32) *CreateDatasetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDatasetResponseBody) SetMessage(v string) *CreateDatasetResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetResponseBody) SetSuccess(v bool) *CreateDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
