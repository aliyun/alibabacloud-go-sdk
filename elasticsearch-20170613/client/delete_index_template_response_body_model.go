// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIndexTemplateResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteIndexTemplateResponseBody
	GetResult() *bool
}

type DeleteIndexTemplateResponseBody struct {
	// example:
	//
	// A0761F7E-0B50-46B9-8CAA-EBB3A420****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteIndexTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIndexTemplateResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteIndexTemplateResponseBody) SetRequestId(v string) *DeleteIndexTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexTemplateResponseBody) SetResult(v bool) *DeleteIndexTemplateResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteIndexTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
