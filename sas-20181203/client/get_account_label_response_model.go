// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountLabelResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountLabelResponseBody) *GetAccountLabelResponse
	GetBody() *GetAccountLabelResponseBody
}

type GetAccountLabelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountLabelResponse) GoString() string {
	return s.String()
}

func (s *GetAccountLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountLabelResponse) GetBody() *GetAccountLabelResponseBody {
	return s.Body
}

func (s *GetAccountLabelResponse) SetHeaders(v map[string]*string) *GetAccountLabelResponse {
	s.Headers = v
	return s
}

func (s *GetAccountLabelResponse) SetStatusCode(v int32) *GetAccountLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountLabelResponse) SetBody(v *GetAccountLabelResponseBody) *GetAccountLabelResponse {
	s.Body = v
	return s
}

func (s *GetAccountLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
