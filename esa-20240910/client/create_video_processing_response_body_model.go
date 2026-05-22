// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoProcessingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateVideoProcessingResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateVideoProcessingResponseBody
	GetRequestId() *string
}

type CreateVideoProcessingResponseBody struct {
	ConfigId  *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVideoProcessingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoProcessingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVideoProcessingResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateVideoProcessingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVideoProcessingResponseBody) SetConfigId(v int64) *CreateVideoProcessingResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateVideoProcessingResponseBody) SetRequestId(v string) *CreateVideoProcessingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVideoProcessingResponseBody) Validate() error {
	return dara.Validate(s)
}
