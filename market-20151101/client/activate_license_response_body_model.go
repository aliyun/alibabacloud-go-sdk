// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ActivateLicenseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ActivateLicenseResponseBody
	GetSuccess() *bool
}

type ActivateLicenseResponseBody struct {
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActivateLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateLicenseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ActivateLicenseResponseBody) SetRequestId(v string) *ActivateLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateLicenseResponseBody) SetSuccess(v bool) *ActivateLicenseResponseBody {
	s.Success = &v
	return s
}

func (s *ActivateLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
