// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterSceneConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCasterSceneConfigResponseBody
	GetRequestId() *string
}

type DeleteCasterSceneConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB9*********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCasterSceneConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterSceneConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCasterSceneConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCasterSceneConfigResponseBody) SetRequestId(v string) *DeleteCasterSceneConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCasterSceneConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
