// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudAccessResponseBody
	GetRequestId() *string
}

type DeleteCloudAccessResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 082FAB35-6AB9-4FD5-8750-D36673548E76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAccessResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudAccessResponseBody) SetRequestId(v string) *DeleteCloudAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
