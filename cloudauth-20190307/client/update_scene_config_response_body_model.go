// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSceneConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSceneConfigResponseBody
	GetRequestId() *string
}

type UpdateSceneConfigResponseBody struct {
	// The ID of this request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSceneConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSceneConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSceneConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSceneConfigResponseBody) SetRequestId(v string) *UpdateSceneConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSceneConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
