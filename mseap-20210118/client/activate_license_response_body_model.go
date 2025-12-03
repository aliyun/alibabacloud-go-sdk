// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ActivateLicenseResponseBody
	GetData() *bool
	SetRequestId(v string) *ActivateLicenseResponseBody
	GetRequestId() *string
}

type ActivateLicenseResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 88EDA98E-6BE7-55DA-9EB6-B6444B882C59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivateLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBody) GetData() *bool {
	return s.Data
}

func (s *ActivateLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateLicenseResponseBody) SetData(v bool) *ActivateLicenseResponseBody {
	s.Data = &v
	return s
}

func (s *ActivateLicenseResponseBody) SetRequestId(v string) *ActivateLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
