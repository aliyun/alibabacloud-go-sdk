// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWatermarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWatermarkResponseBody
	GetRequestId() *string
}

type DeleteWatermarkResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWatermarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWatermarkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWatermarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWatermarkResponseBody) SetRequestId(v string) *DeleteWatermarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWatermarkResponseBody) Validate() error {
	return dara.Validate(s)
}
