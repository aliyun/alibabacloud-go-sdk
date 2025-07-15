// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressCloudConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEccId(v string) *CreateExpressCloudConnectionResponseBody
	GetEccId() *string
	SetRequestId(v string) *CreateExpressCloudConnectionResponseBody
	GetRequestId() *string
}

type CreateExpressCloudConnectionResponseBody struct {
	// The ID of the ECC instance.
	//
	// example:
	//
	// ecc-jg************
	EccId *string `json:"EccId,omitempty" xml:"EccId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C004F022-1CC2-4958-9937-675513A2CD7E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExpressCloudConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressCloudConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExpressCloudConnectionResponseBody) GetEccId() *string {
	return s.EccId
}

func (s *CreateExpressCloudConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExpressCloudConnectionResponseBody) SetEccId(v string) *CreateExpressCloudConnectionResponseBody {
	s.EccId = &v
	return s
}

func (s *CreateExpressCloudConnectionResponseBody) SetRequestId(v string) *CreateExpressCloudConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExpressCloudConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
