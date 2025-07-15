// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveStreamWatermarkResponseBody
	GetRequestId() *string
}

type DeleteLiveStreamWatermarkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5c6a2a0df228-4a64-af62-20e91b******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveStreamWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveStreamWatermarkResponseBody) SetRequestId(v string) *DeleteLiveStreamWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveStreamWatermarkResponseBody) Validate() error {
	return dara.Validate(s)
}
