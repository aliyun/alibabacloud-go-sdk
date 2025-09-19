// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTrailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceTrailResponseBody
	GetRequestId() *string
}

type CreateServiceTrailResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceTrailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTrailResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceTrailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceTrailResponseBody) SetRequestId(v string) *CreateServiceTrailResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceTrailResponseBody) Validate() error {
	return dara.Validate(s)
}
