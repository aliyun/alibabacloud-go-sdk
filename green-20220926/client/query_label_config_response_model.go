// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLabelConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLabelConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLabelConfigResponse
	GetStatusCode() *int32
	SetBody(v *QueryLabelConfigResponseBody) *QueryLabelConfigResponse
	GetBody() *QueryLabelConfigResponseBody
}

type QueryLabelConfigResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLabelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLabelConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLabelConfigResponse) GoString() string {
	return s.String()
}

func (s *QueryLabelConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLabelConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLabelConfigResponse) GetBody() *QueryLabelConfigResponseBody {
	return s.Body
}

func (s *QueryLabelConfigResponse) SetHeaders(v map[string]*string) *QueryLabelConfigResponse {
	s.Headers = v
	return s
}

func (s *QueryLabelConfigResponse) SetStatusCode(v int32) *QueryLabelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLabelConfigResponse) SetBody(v *QueryLabelConfigResponseBody) *QueryLabelConfigResponse {
	s.Body = v
	return s
}

func (s *QueryLabelConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
