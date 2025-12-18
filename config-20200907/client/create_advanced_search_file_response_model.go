// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedSearchFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAdvancedSearchFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAdvancedSearchFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateAdvancedSearchFileResponseBody) *CreateAdvancedSearchFileResponse
	GetBody() *CreateAdvancedSearchFileResponseBody
}

type CreateAdvancedSearchFileResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAdvancedSearchFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAdvancedSearchFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedSearchFileResponse) GoString() string {
	return s.String()
}

func (s *CreateAdvancedSearchFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAdvancedSearchFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAdvancedSearchFileResponse) GetBody() *CreateAdvancedSearchFileResponseBody {
	return s.Body
}

func (s *CreateAdvancedSearchFileResponse) SetHeaders(v map[string]*string) *CreateAdvancedSearchFileResponse {
	s.Headers = v
	return s
}

func (s *CreateAdvancedSearchFileResponse) SetStatusCode(v int32) *CreateAdvancedSearchFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAdvancedSearchFileResponse) SetBody(v *CreateAdvancedSearchFileResponseBody) *CreateAdvancedSearchFileResponse {
	s.Body = v
	return s
}

func (s *CreateAdvancedSearchFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
