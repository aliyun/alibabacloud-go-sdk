// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSiteCustomLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSiteCustomLogResponseBody
	GetRequestId() *string
}

type CreateSiteCustomLogResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSiteCustomLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSiteCustomLogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSiteCustomLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSiteCustomLogResponseBody) SetRequestId(v string) *CreateSiteCustomLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSiteCustomLogResponseBody) Validate() error {
	return dara.Validate(s)
}
