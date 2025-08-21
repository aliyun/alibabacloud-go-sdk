// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudAppResponseBody
	GetRequestId() *string
}

type DeleteCloudAppResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudAppResponseBody) SetRequestId(v string) *DeleteCloudAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudAppResponseBody) Validate() error {
	return dara.Validate(s)
}
