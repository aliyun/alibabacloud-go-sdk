// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppDetailTopoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiAppDetailTopoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiAppDetailTopoResponse
	GetStatusCode() *int32
	SetBody(v *GetAiAppDetailTopoResponseBody) *GetAiAppDetailTopoResponse
	GetBody() *GetAiAppDetailTopoResponseBody
}

type GetAiAppDetailTopoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiAppDetailTopoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiAppDetailTopoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailTopoResponse) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailTopoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiAppDetailTopoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiAppDetailTopoResponse) GetBody() *GetAiAppDetailTopoResponseBody {
	return s.Body
}

func (s *GetAiAppDetailTopoResponse) SetHeaders(v map[string]*string) *GetAiAppDetailTopoResponse {
	s.Headers = v
	return s
}

func (s *GetAiAppDetailTopoResponse) SetStatusCode(v int32) *GetAiAppDetailTopoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiAppDetailTopoResponse) SetBody(v *GetAiAppDetailTopoResponseBody) *GetAiAppDetailTopoResponse {
	s.Body = v
	return s
}

func (s *GetAiAppDetailTopoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
