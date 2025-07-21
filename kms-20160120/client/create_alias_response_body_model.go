// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAliasResponseBody
	GetRequestId() *string
}

type CreateAliasResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 1d2baaf3-d357-46c2-832e-13560c2bd9cd
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAliasResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAliasResponseBody) SetRequestId(v string) *CreateAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAliasResponseBody) Validate() error {
	return dara.Validate(s)
}
