// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKibanaPvlNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKibanaPvlNetworkResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateKibanaPvlNetworkResponseBody
	GetResult() *bool
}

type UpdateKibanaPvlNetworkResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateKibanaPvlNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaPvlNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKibanaPvlNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKibanaPvlNetworkResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateKibanaPvlNetworkResponseBody) SetRequestId(v string) *UpdateKibanaPvlNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKibanaPvlNetworkResponseBody) SetResult(v bool) *UpdateKibanaPvlNetworkResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateKibanaPvlNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
