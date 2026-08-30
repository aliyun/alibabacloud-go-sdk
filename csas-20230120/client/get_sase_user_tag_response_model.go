// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSaseUserTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSaseUserTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSaseUserTagResponse
	GetStatusCode() *int32
	SetBody(v *GetSaseUserTagResponseBody) *GetSaseUserTagResponse
	GetBody() *GetSaseUserTagResponseBody
}

type GetSaseUserTagResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSaseUserTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSaseUserTagResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSaseUserTagResponse) GoString() string {
	return s.String()
}

func (s *GetSaseUserTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSaseUserTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSaseUserTagResponse) GetBody() *GetSaseUserTagResponseBody {
	return s.Body
}

func (s *GetSaseUserTagResponse) SetHeaders(v map[string]*string) *GetSaseUserTagResponse {
	s.Headers = v
	return s
}

func (s *GetSaseUserTagResponse) SetStatusCode(v int32) *GetSaseUserTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSaseUserTagResponse) SetBody(v *GetSaseUserTagResponseBody) *GetSaseUserTagResponse {
	s.Body = v
	return s
}

func (s *GetSaseUserTagResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
