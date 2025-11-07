// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSceneConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSceneConfigResponseBody
	GetRequestId() *string
}

type CreateSceneConfigResponseBody struct {
	// ID of this request.
	//
	// example:
	//
	// B506328A-D84B-4750-82C7-6A207C585CF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSceneConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSceneConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSceneConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSceneConfigResponseBody) SetRequestId(v string) *CreateSceneConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSceneConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
