// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAppSeoIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAppSeoIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAppSeoIndexResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAppSeoIndexResponseBody) *SubmitAppSeoIndexResponse
	GetBody() *SubmitAppSeoIndexResponseBody
}

type SubmitAppSeoIndexResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAppSeoIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAppSeoIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAppSeoIndexResponse) GoString() string {
	return s.String()
}

func (s *SubmitAppSeoIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAppSeoIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAppSeoIndexResponse) GetBody() *SubmitAppSeoIndexResponseBody {
	return s.Body
}

func (s *SubmitAppSeoIndexResponse) SetHeaders(v map[string]*string) *SubmitAppSeoIndexResponse {
	s.Headers = v
	return s
}

func (s *SubmitAppSeoIndexResponse) SetStatusCode(v int32) *SubmitAppSeoIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAppSeoIndexResponse) SetBody(v *SubmitAppSeoIndexResponseBody) *SubmitAppSeoIndexResponse {
	s.Body = v
	return s
}

func (s *SubmitAppSeoIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
