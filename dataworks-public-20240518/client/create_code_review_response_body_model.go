// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeReviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateCodeReviewResponseBody
	GetId() *string
	SetRequestId(v string) *CreateCodeReviewResponseBody
	GetRequestId() *string
}

type CreateCodeReviewResponseBody struct {
	// The unique identifier of the created code review.
	//
	// example:
	//
	// review-id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCodeReviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeReviewResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCodeReviewResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateCodeReviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCodeReviewResponseBody) SetId(v string) *CreateCodeReviewResponseBody {
	s.Id = &v
	return s
}

func (s *CreateCodeReviewResponseBody) SetRequestId(v string) *CreateCodeReviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCodeReviewResponseBody) Validate() error {
	return dara.Validate(s)
}
