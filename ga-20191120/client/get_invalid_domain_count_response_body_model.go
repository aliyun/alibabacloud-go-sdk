// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInvalidDomainCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvalidDomainCount(v string) *GetInvalidDomainCountResponseBody
	GetInvalidDomainCount() *string
	SetRequestId(v string) *GetInvalidDomainCountResponseBody
	GetRequestId() *string
}

type GetInvalidDomainCountResponseBody struct {
	// The number of invalid domain names.
	//
	// example:
	//
	// 1
	InvalidDomainCount *string `json:"InvalidDomainCount,omitempty" xml:"InvalidDomainCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInvalidDomainCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInvalidDomainCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetInvalidDomainCountResponseBody) GetInvalidDomainCount() *string {
	return s.InvalidDomainCount
}

func (s *GetInvalidDomainCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInvalidDomainCountResponseBody) SetInvalidDomainCount(v string) *GetInvalidDomainCountResponseBody {
	s.InvalidDomainCount = &v
	return s
}

func (s *GetInvalidDomainCountResponseBody) SetRequestId(v string) *GetInvalidDomainCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInvalidDomainCountResponseBody) Validate() error {
	return dara.Validate(s)
}
