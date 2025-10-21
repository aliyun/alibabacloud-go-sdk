// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavepointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSavepointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSavepointResponse
	GetStatusCode() *int32
	SetBody(v *GetSavepointResponseBody) *GetSavepointResponse
	GetBody() *GetSavepointResponseBody
}

type GetSavepointResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSavepointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSavepointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSavepointResponse) GoString() string {
	return s.String()
}

func (s *GetSavepointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSavepointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSavepointResponse) GetBody() *GetSavepointResponseBody {
	return s.Body
}

func (s *GetSavepointResponse) SetHeaders(v map[string]*string) *GetSavepointResponse {
	s.Headers = v
	return s
}

func (s *GetSavepointResponse) SetStatusCode(v int32) *GetSavepointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavepointResponse) SetBody(v *GetSavepointResponseBody) *GetSavepointResponse {
	s.Body = v
	return s
}

func (s *GetSavepointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
