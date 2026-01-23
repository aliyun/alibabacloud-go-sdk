// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualitySchedulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteQualitySchedulesResponseBody
	GetCode() *string
	SetData(v int32) *DeleteQualitySchedulesResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *DeleteQualitySchedulesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteQualitySchedulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteQualitySchedulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQualitySchedulesResponseBody
	GetSuccess() *bool
}

type DeleteQualitySchedulesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQualitySchedulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualitySchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualitySchedulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteQualitySchedulesResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteQualitySchedulesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteQualitySchedulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteQualitySchedulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQualitySchedulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQualitySchedulesResponseBody) SetCode(v string) *DeleteQualitySchedulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQualitySchedulesResponseBody) SetData(v int32) *DeleteQualitySchedulesResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualitySchedulesResponseBody) SetHttpStatusCode(v int32) *DeleteQualitySchedulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteQualitySchedulesResponseBody) SetMessage(v string) *DeleteQualitySchedulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQualitySchedulesResponseBody) SetRequestId(v string) *DeleteQualitySchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualitySchedulesResponseBody) SetSuccess(v bool) *DeleteQualitySchedulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQualitySchedulesResponseBody) Validate() error {
	return dara.Validate(s)
}
