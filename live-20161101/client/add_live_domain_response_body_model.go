// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLiveDomainResponseBody
	GetRequestId() *string
}

type AddLiveDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLiveDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLiveDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AddLiveDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLiveDomainResponseBody) SetRequestId(v string) *AddLiveDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLiveDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
