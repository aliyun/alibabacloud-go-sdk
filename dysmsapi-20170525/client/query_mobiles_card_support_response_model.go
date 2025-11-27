// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMobilesCardSupportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMobilesCardSupportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMobilesCardSupportResponse
	GetStatusCode() *int32
	SetBody(v *QueryMobilesCardSupportResponseBody) *QueryMobilesCardSupportResponse
	GetBody() *QueryMobilesCardSupportResponseBody
}

type QueryMobilesCardSupportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMobilesCardSupportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMobilesCardSupportResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMobilesCardSupportResponse) GoString() string {
	return s.String()
}

func (s *QueryMobilesCardSupportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMobilesCardSupportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMobilesCardSupportResponse) GetBody() *QueryMobilesCardSupportResponseBody {
	return s.Body
}

func (s *QueryMobilesCardSupportResponse) SetHeaders(v map[string]*string) *QueryMobilesCardSupportResponse {
	s.Headers = v
	return s
}

func (s *QueryMobilesCardSupportResponse) SetStatusCode(v int32) *QueryMobilesCardSupportResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMobilesCardSupportResponse) SetBody(v *QueryMobilesCardSupportResponseBody) *QueryMobilesCardSupportResponse {
	s.Body = v
	return s
}

func (s *QueryMobilesCardSupportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
