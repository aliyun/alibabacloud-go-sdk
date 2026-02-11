// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDatasetResponse
	GetStatusCode() *int32
	SetBody(v *QueryDatasetResponseBody) *QueryDatasetResponse
	GetBody() *QueryDatasetResponseBody
}

type QueryDatasetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetResponse) GoString() string {
	return s.String()
}

func (s *QueryDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDatasetResponse) GetBody() *QueryDatasetResponseBody {
	return s.Body
}

func (s *QueryDatasetResponse) SetHeaders(v map[string]*string) *QueryDatasetResponse {
	s.Headers = v
	return s
}

func (s *QueryDatasetResponse) SetStatusCode(v int32) *QueryDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDatasetResponse) SetBody(v *QueryDatasetResponseBody) *QueryDatasetResponse {
	s.Body = v
	return s
}

func (s *QueryDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
