// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteJobTemplateResponseBody
	GetRequestId() *string
}

type DeleteJobTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteJobTemplateResponseBody) SetRequestId(v string) *DeleteJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteJobTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
