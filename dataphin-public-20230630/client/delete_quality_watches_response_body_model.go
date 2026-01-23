// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityWatchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteQualityWatchesResponseBody
	GetCode() *string
	SetData(v int32) *DeleteQualityWatchesResponseBody
	GetData() *int32
	SetHttpStatusCode(v int32) *DeleteQualityWatchesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteQualityWatchesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteQualityWatchesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQualityWatchesResponseBody
	GetSuccess() *bool
}

type DeleteQualityWatchesResponseBody struct {
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

func (s DeleteQualityWatchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityWatchesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityWatchesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteQualityWatchesResponseBody) GetData() *int32 {
	return s.Data
}

func (s *DeleteQualityWatchesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteQualityWatchesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteQualityWatchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQualityWatchesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQualityWatchesResponseBody) SetCode(v string) *DeleteQualityWatchesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQualityWatchesResponseBody) SetData(v int32) *DeleteQualityWatchesResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualityWatchesResponseBody) SetHttpStatusCode(v int32) *DeleteQualityWatchesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteQualityWatchesResponseBody) SetMessage(v string) *DeleteQualityWatchesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQualityWatchesResponseBody) SetRequestId(v string) *DeleteQualityWatchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityWatchesResponseBody) SetSuccess(v bool) *DeleteQualityWatchesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQualityWatchesResponseBody) Validate() error {
	return dara.Validate(s)
}
