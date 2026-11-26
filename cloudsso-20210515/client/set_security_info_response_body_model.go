// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSecurityInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSecurityInfoResponseBody
	GetRequestId() *string
}

type SetSecurityInfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F44F02EC-70D1-5E51-8E8E-FA9AC4EF952A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSecurityInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSecurityInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SetSecurityInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSecurityInfoResponseBody) SetRequestId(v string) *SetSecurityInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSecurityInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
