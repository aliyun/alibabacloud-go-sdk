// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocContentResponse
	GetStatusCode() *int32
	SetBody(v *GetDocContentResponseBody) *GetDocContentResponse
	GetBody() *GetDocContentResponseBody
}

type GetDocContentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocContentResponse) GoString() string {
	return s.String()
}

func (s *GetDocContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocContentResponse) GetBody() *GetDocContentResponseBody {
	return s.Body
}

func (s *GetDocContentResponse) SetHeaders(v map[string]*string) *GetDocContentResponse {
	s.Headers = v
	return s
}

func (s *GetDocContentResponse) SetStatusCode(v int32) *GetDocContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocContentResponse) SetBody(v *GetDocContentResponseBody) *GetDocContentResponse {
	s.Body = v
	return s
}

func (s *GetDocContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
