// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateInstanceAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceAccountResponseBody
	GetSuccess() *bool
}

type CreateInstanceAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// FF13E47D-4E38-4A5A-BA68-4E610EVF56DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceAccountResponseBody) SetRequestId(v string) *CreateInstanceAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) SetSuccess(v bool) *CreateInstanceAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
