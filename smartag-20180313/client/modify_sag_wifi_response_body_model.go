// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagWifiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagWifiResponseBody
	GetRequestId() *string
}

type ModifySagWifiResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1120228A-E5E1-4E9C-B56D-96887E1A2B2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagWifiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagWifiResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagWifiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagWifiResponseBody) SetRequestId(v string) *ModifySagWifiResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagWifiResponseBody) Validate() error {
	return dara.Validate(s)
}
