// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageTransformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateImageTransformResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateImageTransformResponseBody
	GetRequestId() *string
}

type CreateImageTransformResponseBody struct {
	// Configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageTransformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageTransformResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageTransformResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateImageTransformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageTransformResponseBody) SetConfigId(v int64) *CreateImageTransformResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateImageTransformResponseBody) SetRequestId(v string) *CreateImageTransformResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageTransformResponseBody) Validate() error {
	return dara.Validate(s)
}
