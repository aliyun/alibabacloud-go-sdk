// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataServiceApiAuthorityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataServiceApiAuthorityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataServiceApiAuthorityResponseBody
	GetSuccess() *bool
}

type DeleteDataServiceApiAuthorityResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataServiceApiAuthorityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataServiceApiAuthorityResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataServiceApiAuthorityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataServiceApiAuthorityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataServiceApiAuthorityResponseBody) SetRequestId(v string) *DeleteDataServiceApiAuthorityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataServiceApiAuthorityResponseBody) SetSuccess(v bool) *DeleteDataServiceApiAuthorityResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataServiceApiAuthorityResponseBody) Validate() error {
	return dara.Validate(s)
}
