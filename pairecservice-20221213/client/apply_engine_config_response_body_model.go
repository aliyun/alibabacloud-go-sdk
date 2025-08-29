// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyEngineConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyEngineConfigResponseBody
	GetRequestId() *string
}

type ApplyEngineConfigResponseBody struct {
	// example:
	//
	// F8F613A9-DF1C-551A-88E1-397A3981A785
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyEngineConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyEngineConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyEngineConfigResponseBody) SetRequestId(v string) *ApplyEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyEngineConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
