// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodSpecificConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVodSpecificConfigResponseBody
	GetRequestId() *string
}

type DeleteVodSpecificConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-****-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVodSpecificConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodSpecificConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVodSpecificConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVodSpecificConfigResponseBody) SetRequestId(v string) *DeleteVodSpecificConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVodSpecificConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
