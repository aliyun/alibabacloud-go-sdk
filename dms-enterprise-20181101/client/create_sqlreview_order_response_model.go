// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSQLReviewOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSQLReviewOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSQLReviewOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateSQLReviewOrderResponseBody) *CreateSQLReviewOrderResponse
	GetBody() *CreateSQLReviewOrderResponseBody
}

type CreateSQLReviewOrderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSQLReviewOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSQLReviewOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLReviewOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateSQLReviewOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSQLReviewOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSQLReviewOrderResponse) GetBody() *CreateSQLReviewOrderResponseBody {
	return s.Body
}

func (s *CreateSQLReviewOrderResponse) SetHeaders(v map[string]*string) *CreateSQLReviewOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateSQLReviewOrderResponse) SetStatusCode(v int32) *CreateSQLReviewOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSQLReviewOrderResponse) SetBody(v *CreateSQLReviewOrderResponseBody) *CreateSQLReviewOrderResponse {
	s.Body = v
	return s
}

func (s *CreateSQLReviewOrderResponse) Validate() error {
	return dara.Validate(s)
}
