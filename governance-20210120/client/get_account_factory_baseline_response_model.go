// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountFactoryBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountFactoryBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountFactoryBaselineResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountFactoryBaselineResponseBody) *GetAccountFactoryBaselineResponse
	GetBody() *GetAccountFactoryBaselineResponseBody
}

type GetAccountFactoryBaselineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountFactoryBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountFactoryBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountFactoryBaselineResponse) GoString() string {
	return s.String()
}

func (s *GetAccountFactoryBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountFactoryBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountFactoryBaselineResponse) GetBody() *GetAccountFactoryBaselineResponseBody {
	return s.Body
}

func (s *GetAccountFactoryBaselineResponse) SetHeaders(v map[string]*string) *GetAccountFactoryBaselineResponse {
	s.Headers = v
	return s
}

func (s *GetAccountFactoryBaselineResponse) SetStatusCode(v int32) *GetAccountFactoryBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountFactoryBaselineResponse) SetBody(v *GetAccountFactoryBaselineResponseBody) *GetAccountFactoryBaselineResponse {
	s.Body = v
	return s
}

func (s *GetAccountFactoryBaselineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
