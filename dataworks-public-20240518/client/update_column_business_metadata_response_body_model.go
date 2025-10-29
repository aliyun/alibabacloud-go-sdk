// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateColumnBusinessMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateColumnBusinessMetadataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateColumnBusinessMetadataResponseBody
	GetSuccess() *bool
}

type UpdateColumnBusinessMetadataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D1E2E5BC-xxxx-xxxx-xxxx-xxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateColumnBusinessMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateColumnBusinessMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateColumnBusinessMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateColumnBusinessMetadataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateColumnBusinessMetadataResponseBody) SetRequestId(v string) *UpdateColumnBusinessMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateColumnBusinessMetadataResponseBody) SetSuccess(v bool) *UpdateColumnBusinessMetadataResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateColumnBusinessMetadataResponseBody) Validate() error {
	return dara.Validate(s)
}
