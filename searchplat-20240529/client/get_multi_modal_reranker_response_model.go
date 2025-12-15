// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiModalRerankerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiModalRerankerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiModalRerankerResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiModalRerankerResponseBody) *GetMultiModalRerankerResponse
	GetBody() *GetMultiModalRerankerResponseBody
}

type GetMultiModalRerankerResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiModalRerankerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiModalRerankerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalRerankerResponse) GoString() string {
	return s.String()
}

func (s *GetMultiModalRerankerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiModalRerankerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiModalRerankerResponse) GetBody() *GetMultiModalRerankerResponseBody {
	return s.Body
}

func (s *GetMultiModalRerankerResponse) SetHeaders(v map[string]*string) *GetMultiModalRerankerResponse {
	s.Headers = v
	return s
}

func (s *GetMultiModalRerankerResponse) SetStatusCode(v int32) *GetMultiModalRerankerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiModalRerankerResponse) SetBody(v *GetMultiModalRerankerResponseBody) *GetMultiModalRerankerResponse {
	s.Body = v
	return s
}

func (s *GetMultiModalRerankerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
