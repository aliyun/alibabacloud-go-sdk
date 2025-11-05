// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteCdnDomainConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteCdnDomainConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteCdnDomainConfigResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteCdnDomainConfigResponseBody) *BatchDeleteCdnDomainConfigResponse
	GetBody() *BatchDeleteCdnDomainConfigResponseBody
}

type BatchDeleteCdnDomainConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteCdnDomainConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteCdnDomainConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteCdnDomainConfigResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteCdnDomainConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteCdnDomainConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteCdnDomainConfigResponse) GetBody() *BatchDeleteCdnDomainConfigResponseBody {
	return s.Body
}

func (s *BatchDeleteCdnDomainConfigResponse) SetHeaders(v map[string]*string) *BatchDeleteCdnDomainConfigResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteCdnDomainConfigResponse) SetStatusCode(v int32) *BatchDeleteCdnDomainConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteCdnDomainConfigResponse) SetBody(v *BatchDeleteCdnDomainConfigResponseBody) *BatchDeleteCdnDomainConfigResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteCdnDomainConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
