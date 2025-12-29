// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicInfoQAResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBasicInfoQAResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBasicInfoQAResponse
	GetStatusCode() *int32
	SetBody(v *GetBasicInfoQAResponseBody) *GetBasicInfoQAResponse
	GetBody() *GetBasicInfoQAResponseBody
}

type GetBasicInfoQAResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBasicInfoQAResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBasicInfoQAResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBasicInfoQAResponse) GoString() string {
	return s.String()
}

func (s *GetBasicInfoQAResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBasicInfoQAResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBasicInfoQAResponse) GetBody() *GetBasicInfoQAResponseBody {
	return s.Body
}

func (s *GetBasicInfoQAResponse) SetHeaders(v map[string]*string) *GetBasicInfoQAResponse {
	s.Headers = v
	return s
}

func (s *GetBasicInfoQAResponse) SetStatusCode(v int32) *GetBasicInfoQAResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicInfoQAResponse) SetBody(v *GetBasicInfoQAResponseBody) *GetBasicInfoQAResponse {
	s.Body = v
	return s
}

func (s *GetBasicInfoQAResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
