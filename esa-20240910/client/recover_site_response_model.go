// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RecoverSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RecoverSiteResponse
	GetStatusCode() *int32
	SetBody(v *RecoverSiteResponseBody) *RecoverSiteResponse
	GetBody() *RecoverSiteResponseBody
}

type RecoverSiteResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s RecoverSiteResponse) GoString() string {
	return s.String()
}

func (s *RecoverSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RecoverSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RecoverSiteResponse) GetBody() *RecoverSiteResponseBody {
	return s.Body
}

func (s *RecoverSiteResponse) SetHeaders(v map[string]*string) *RecoverSiteResponse {
	s.Headers = v
	return s
}

func (s *RecoverSiteResponse) SetStatusCode(v int32) *RecoverSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverSiteResponse) SetBody(v *RecoverSiteResponseBody) *RecoverSiteResponse {
	s.Body = v
	return s
}

func (s *RecoverSiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
