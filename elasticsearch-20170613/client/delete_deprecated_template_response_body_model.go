// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeprecatedTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDeprecatedTemplateResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteDeprecatedTemplateResponseBody
	GetResult() *bool
}

type DeleteDeprecatedTemplateResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteDeprecatedTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeprecatedTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDeprecatedTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDeprecatedTemplateResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteDeprecatedTemplateResponseBody) SetRequestId(v string) *DeleteDeprecatedTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDeprecatedTemplateResponseBody) SetResult(v bool) *DeleteDeprecatedTemplateResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteDeprecatedTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
