// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetL2OssKeyConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetL2OssKeyConfigResponseBody
	GetRequestId() *string
}

type SetL2OssKeyConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetL2OssKeyConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetL2OssKeyConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetL2OssKeyConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetL2OssKeyConfigResponseBody) SetRequestId(v string) *SetL2OssKeyConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetL2OssKeyConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
