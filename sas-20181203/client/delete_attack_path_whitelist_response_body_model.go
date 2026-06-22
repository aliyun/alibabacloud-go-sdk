// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttackPathWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAttackPathWhitelistResponseBody
	GetRequestId() *string
}

type DeleteAttackPathWhitelistResponseBody struct {
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 898F7AA7-CECD-5EC7-AF4D-664C601B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAttackPathWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttackPathWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAttackPathWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAttackPathWhitelistResponseBody) SetRequestId(v string) *DeleteAttackPathWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAttackPathWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
