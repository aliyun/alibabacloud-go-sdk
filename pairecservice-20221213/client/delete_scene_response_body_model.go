// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSceneResponseBody
	GetRequestId() *string
}

type DeleteSceneResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// D75C43DC-3D3A-5CC8-9AAC-8C77306C433B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSceneResponseBody) SetRequestId(v string) *DeleteSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
