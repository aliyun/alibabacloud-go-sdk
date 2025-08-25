// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceVideoTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFaceVideoTemplateResponseBody
	GetRequestId() *string
}

type DeleteFaceVideoTemplateResponseBody struct {
	// example:
	//
	// 2337D184-CD63-57B5-96A2-B60AABCB7B58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaceVideoTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceVideoTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaceVideoTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaceVideoTemplateResponseBody) SetRequestId(v string) *DeleteFaceVideoTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaceVideoTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
