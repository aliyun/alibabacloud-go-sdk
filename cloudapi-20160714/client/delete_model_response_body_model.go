// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteModelResponseBody
	GetRequestId() *string
}

type DeleteModelResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4173F95B-360C-460C-9F6C-4A960B904411
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelResponseBody) SetRequestId(v string) *DeleteModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelResponseBody) Validate() error {
	return dara.Validate(s)
}
