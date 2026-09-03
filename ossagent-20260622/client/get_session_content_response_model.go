// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSessionContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSessionContentResponse
	GetStatusCode() *int32
	SetBody(v *GetSessionContentResponseBody) *GetSessionContentResponse
	GetBody() *GetSessionContentResponseBody
}

type GetSessionContentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSessionContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSessionContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSessionContentResponse) GoString() string {
	return s.String()
}

func (s *GetSessionContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSessionContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSessionContentResponse) GetBody() *GetSessionContentResponseBody {
	return s.Body
}

func (s *GetSessionContentResponse) SetHeaders(v map[string]*string) *GetSessionContentResponse {
	s.Headers = v
	return s
}

func (s *GetSessionContentResponse) SetStatusCode(v int32) *GetSessionContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSessionContentResponse) SetBody(v *GetSessionContentResponseBody) *GetSessionContentResponse {
	s.Body = v
	return s
}

func (s *GetSessionContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
