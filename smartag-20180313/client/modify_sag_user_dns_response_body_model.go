// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySagUserDnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySagUserDnsResponseBody
	GetRequestId() *string
}

type ModifySagUserDnsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1120228A-E5E1-4E9C-B56D-96887E1A2B2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySagUserDnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySagUserDnsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySagUserDnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySagUserDnsResponseBody) SetRequestId(v string) *ModifySagUserDnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySagUserDnsResponseBody) Validate() error {
	return dara.Validate(s)
}
