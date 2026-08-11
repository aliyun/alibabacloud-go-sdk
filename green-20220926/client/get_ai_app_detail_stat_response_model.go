// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppDetailStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAiAppDetailStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAiAppDetailStatResponse
	GetStatusCode() *int32
	SetBody(v *GetAiAppDetailStatResponseBody) *GetAiAppDetailStatResponse
	GetBody() *GetAiAppDetailStatResponseBody
}

type GetAiAppDetailStatResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAiAppDetailStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAiAppDetailStatResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailStatResponse) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAiAppDetailStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAiAppDetailStatResponse) GetBody() *GetAiAppDetailStatResponseBody {
	return s.Body
}

func (s *GetAiAppDetailStatResponse) SetHeaders(v map[string]*string) *GetAiAppDetailStatResponse {
	s.Headers = v
	return s
}

func (s *GetAiAppDetailStatResponse) SetStatusCode(v int32) *GetAiAppDetailStatResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAiAppDetailStatResponse) SetBody(v *GetAiAppDetailStatResponseBody) *GetAiAppDetailStatResponse {
	s.Body = v
	return s
}

func (s *GetAiAppDetailStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
