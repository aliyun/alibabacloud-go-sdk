// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateCdnUrlAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateCdnUrlAuthResponseBody
	GetRequestId() *string
	SetResult(v string) *ValidateCdnUrlAuthResponseBody
	GetResult() *string
}

type ValidateCdnUrlAuthResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ValidateCdnUrlAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateCdnUrlAuthResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateCdnUrlAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateCdnUrlAuthResponseBody) GetResult() *string {
	return s.Result
}

func (s *ValidateCdnUrlAuthResponseBody) SetRequestId(v string) *ValidateCdnUrlAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateCdnUrlAuthResponseBody) SetResult(v string) *ValidateCdnUrlAuthResponseBody {
	s.Result = &v
	return s
}

func (s *ValidateCdnUrlAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
