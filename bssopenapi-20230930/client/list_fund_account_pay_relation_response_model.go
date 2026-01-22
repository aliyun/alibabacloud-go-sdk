// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFundAccountPayRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFundAccountPayRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFundAccountPayRelationResponse
	GetStatusCode() *int32
	SetBody(v *ListFundAccountPayRelationResponseBody) *ListFundAccountPayRelationResponse
	GetBody() *ListFundAccountPayRelationResponseBody
}

type ListFundAccountPayRelationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFundAccountPayRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFundAccountPayRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFundAccountPayRelationResponse) GoString() string {
	return s.String()
}

func (s *ListFundAccountPayRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFundAccountPayRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFundAccountPayRelationResponse) GetBody() *ListFundAccountPayRelationResponseBody {
	return s.Body
}

func (s *ListFundAccountPayRelationResponse) SetHeaders(v map[string]*string) *ListFundAccountPayRelationResponse {
	s.Headers = v
	return s
}

func (s *ListFundAccountPayRelationResponse) SetStatusCode(v int32) *ListFundAccountPayRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFundAccountPayRelationResponse) SetBody(v *ListFundAccountPayRelationResponseBody) *ListFundAccountPayRelationResponse {
	s.Body = v
	return s
}

func (s *ListFundAccountPayRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
