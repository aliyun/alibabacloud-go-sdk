// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamLogContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStreamLogContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStreamLogContentResponse
	GetStatusCode() *int32
	SetBody(v *GetStreamLogContentResponseBody) *GetStreamLogContentResponse
	GetBody() *GetStreamLogContentResponseBody
}

type GetStreamLogContentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStreamLogContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStreamLogContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStreamLogContentResponse) GoString() string {
	return s.String()
}

func (s *GetStreamLogContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStreamLogContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStreamLogContentResponse) GetBody() *GetStreamLogContentResponseBody {
	return s.Body
}

func (s *GetStreamLogContentResponse) SetHeaders(v map[string]*string) *GetStreamLogContentResponse {
	s.Headers = v
	return s
}

func (s *GetStreamLogContentResponse) SetStatusCode(v int32) *GetStreamLogContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStreamLogContentResponse) SetBody(v *GetStreamLogContentResponseBody) *GetStreamLogContentResponse {
	s.Body = v
	return s
}

func (s *GetStreamLogContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
