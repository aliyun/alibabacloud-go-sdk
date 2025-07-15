// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLiveDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartLiveDomainResponseBody
	GetRequestId() *string
}

type StartLiveDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartLiveDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartLiveDomainResponseBody) GoString() string {
	return s.String()
}

func (s *StartLiveDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartLiveDomainResponseBody) SetRequestId(v string) *StartLiveDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartLiveDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
