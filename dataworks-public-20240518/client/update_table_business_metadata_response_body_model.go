// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableBusinessMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTableBusinessMetadataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateTableBusinessMetadataResponseBody
	GetSuccess() *bool
}

type UpdateTableBusinessMetadataResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 7C352CB7-CD88-XXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTableBusinessMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableBusinessMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableBusinessMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTableBusinessMetadataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateTableBusinessMetadataResponseBody) SetRequestId(v string) *UpdateTableBusinessMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableBusinessMetadataResponseBody) SetSuccess(v bool) *UpdateTableBusinessMetadataResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTableBusinessMetadataResponseBody) Validate() error {
	return dara.Validate(s)
}
