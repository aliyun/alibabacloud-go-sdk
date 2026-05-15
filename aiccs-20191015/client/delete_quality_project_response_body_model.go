// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteQualityProjectResponseBody
	GetCode() *string
	SetData(v string) *DeleteQualityProjectResponseBody
	GetData() *string
	SetMessage(v string) *DeleteQualityProjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteQualityProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteQualityProjectResponseBody
	GetSuccess() *bool
}

type DeleteQualityProjectResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteQualityProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQualityProjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteQualityProjectResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteQualityProjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteQualityProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQualityProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteQualityProjectResponseBody) SetCode(v string) *DeleteQualityProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteQualityProjectResponseBody) SetData(v string) *DeleteQualityProjectResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteQualityProjectResponseBody) SetMessage(v string) *DeleteQualityProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteQualityProjectResponseBody) SetRequestId(v string) *DeleteQualityProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQualityProjectResponseBody) SetSuccess(v bool) *DeleteQualityProjectResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteQualityProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
