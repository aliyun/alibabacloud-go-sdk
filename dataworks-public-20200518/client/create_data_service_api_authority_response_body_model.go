// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceApiAuthorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDataServiceApiAuthorityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataServiceApiAuthorityResponseBody
	GetSuccess() *bool
}

type CreateDataServiceApiAuthorityResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the authorization was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataServiceApiAuthorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceApiAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataServiceApiAuthorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataServiceApiAuthorityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataServiceApiAuthorityResponseBody) SetRequestId(v string) *CreateDataServiceApiAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataServiceApiAuthorityResponseBody) SetSuccess(v bool) *CreateDataServiceApiAuthorityResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataServiceApiAuthorityResponseBody) Validate() error {
	return dara.Validate(s)
}
