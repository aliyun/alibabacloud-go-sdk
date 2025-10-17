// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkTemplateFullTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSparkTemplateFullTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSparkTemplateFullTreeResponse
	GetStatusCode() *int32
	SetBody(v *GetSparkTemplateFullTreeResponseBody) *GetSparkTemplateFullTreeResponse
	GetBody() *GetSparkTemplateFullTreeResponseBody
}

type GetSparkTemplateFullTreeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSparkTemplateFullTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSparkTemplateFullTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSparkTemplateFullTreeResponse) GoString() string {
	return s.String()
}

func (s *GetSparkTemplateFullTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSparkTemplateFullTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSparkTemplateFullTreeResponse) GetBody() *GetSparkTemplateFullTreeResponseBody {
	return s.Body
}

func (s *GetSparkTemplateFullTreeResponse) SetHeaders(v map[string]*string) *GetSparkTemplateFullTreeResponse {
	s.Headers = v
	return s
}

func (s *GetSparkTemplateFullTreeResponse) SetStatusCode(v int32) *GetSparkTemplateFullTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSparkTemplateFullTreeResponse) SetBody(v *GetSparkTemplateFullTreeResponseBody) *GetSparkTemplateFullTreeResponse {
	s.Body = v
	return s
}

func (s *GetSparkTemplateFullTreeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
