// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSceneConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSceneConfigResponseBody
	GetRequestId() *string
}

type DeleteSceneConfigResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSceneConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSceneConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSceneConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSceneConfigResponseBody) SetRequestId(v string) *DeleteSceneConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSceneConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
