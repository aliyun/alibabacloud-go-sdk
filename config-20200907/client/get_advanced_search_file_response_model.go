// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdvancedSearchFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAdvancedSearchFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAdvancedSearchFileResponse
	GetStatusCode() *int32
	SetBody(v *GetAdvancedSearchFileResponseBody) *GetAdvancedSearchFileResponse
	GetBody() *GetAdvancedSearchFileResponseBody
}

type GetAdvancedSearchFileResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAdvancedSearchFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAdvancedSearchFileResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAdvancedSearchFileResponse) GoString() string {
	return s.String()
}

func (s *GetAdvancedSearchFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAdvancedSearchFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAdvancedSearchFileResponse) GetBody() *GetAdvancedSearchFileResponseBody {
	return s.Body
}

func (s *GetAdvancedSearchFileResponse) SetHeaders(v map[string]*string) *GetAdvancedSearchFileResponse {
	s.Headers = v
	return s
}

func (s *GetAdvancedSearchFileResponse) SetStatusCode(v int32) *GetAdvancedSearchFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAdvancedSearchFileResponse) SetBody(v *GetAdvancedSearchFileResponseBody) *GetAdvancedSearchFileResponse {
	s.Body = v
	return s
}

func (s *GetAdvancedSearchFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
