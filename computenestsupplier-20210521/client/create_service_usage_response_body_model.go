// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceUsageResponseBody
	GetRequestId() *string
}

type CreateServiceUsageResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// E73F09DC-6C13-5CB1-A10F-7A4E125ABD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceUsageResponseBody) SetRequestId(v string) *CreateServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
