// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoProcessingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVideoProcessingResponseBody
	GetRequestId() *string
}

type UpdateVideoProcessingResponseBody struct {
	// example:
	//
	// 3558df77-8a7a-4060-a900-2d7949403836
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVideoProcessingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoProcessingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVideoProcessingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVideoProcessingResponseBody) SetRequestId(v string) *UpdateVideoProcessingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVideoProcessingResponseBody) Validate() error {
	return dara.Validate(s)
}
