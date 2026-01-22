// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFundAccountPayRelationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFundAccountPayRelationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFundAccountPayRelationResponse
	GetStatusCode() *int32
	SetBody(v *CreateFundAccountPayRelationResponseBody) *CreateFundAccountPayRelationResponse
	GetBody() *CreateFundAccountPayRelationResponseBody
}

type CreateFundAccountPayRelationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFundAccountPayRelationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFundAccountPayRelationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFundAccountPayRelationResponse) GoString() string {
	return s.String()
}

func (s *CreateFundAccountPayRelationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFundAccountPayRelationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFundAccountPayRelationResponse) GetBody() *CreateFundAccountPayRelationResponseBody {
	return s.Body
}

func (s *CreateFundAccountPayRelationResponse) SetHeaders(v map[string]*string) *CreateFundAccountPayRelationResponse {
	s.Headers = v
	return s
}

func (s *CreateFundAccountPayRelationResponse) SetStatusCode(v int32) *CreateFundAccountPayRelationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFundAccountPayRelationResponse) SetBody(v *CreateFundAccountPayRelationResponseBody) *CreateFundAccountPayRelationResponse {
	s.Body = v
	return s
}

func (s *CreateFundAccountPayRelationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
