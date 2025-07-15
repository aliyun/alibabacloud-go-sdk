// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAppKeyResponseBody
	GetRequestId() *string
}

type CreateAppKeyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5D524309-6BED-5BB4-A735-F7D9F98B7B88
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppKeyResponseBody) SetRequestId(v string) *CreateAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
