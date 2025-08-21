// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateIndexTemplateResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateIndexTemplateResponseBody
	GetResult() *string
}

type CreateIndexTemplateResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// index-template
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateIndexTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIndexTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIndexTemplateResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateIndexTemplateResponseBody) SetRequestId(v string) *CreateIndexTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIndexTemplateResponseBody) SetResult(v string) *CreateIndexTemplateResponseBody {
	s.Result = &v
	return s
}

func (s *CreateIndexTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
