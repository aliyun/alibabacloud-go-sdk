// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCodeReviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCodeReviewResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCodeReviewResponseBody
	GetSuccess() *bool
}

type UpdateCodeReviewResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// request-id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the approval is successful. This does not indicate that the publish process is complete.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCodeReviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCodeReviewResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCodeReviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCodeReviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCodeReviewResponseBody) SetRequestId(v string) *UpdateCodeReviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCodeReviewResponseBody) SetSuccess(v bool) *UpdateCodeReviewResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCodeReviewResponseBody) Validate() error {
	return dara.Validate(s)
}
