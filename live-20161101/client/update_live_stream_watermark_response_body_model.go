// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveStreamWatermarkResponseBody
	GetRequestId() *string
}

type UpdateLiveStreamWatermarkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5c6a2a0df228-4a64-af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveStreamWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveStreamWatermarkResponseBody) SetRequestId(v string) *UpdateLiveStreamWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveStreamWatermarkResponseBody) Validate() error {
	return dara.Validate(s)
}
