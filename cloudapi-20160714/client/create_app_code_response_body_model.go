// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAppCodeResponseBody
	GetRequestId() *string
}

type CreateAppCodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 418DAAE7-A0C2-5E9C-ADFF-4CD14A474F88
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppCodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppCodeResponseBody) SetRequestId(v string) *CreateAppCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
