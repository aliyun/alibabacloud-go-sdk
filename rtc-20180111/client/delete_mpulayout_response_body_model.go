// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMPULayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMPULayoutResponseBody
	GetRequestId() *string
}

type DeleteMPULayoutResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMPULayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMPULayoutResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMPULayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMPULayoutResponseBody) SetRequestId(v string) *DeleteMPULayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMPULayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
