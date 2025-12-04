// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDatasetResponse
	GetStatusCode() *int32
	SetBody(v *GetDatasetResponseBody) *GetDatasetResponse
	GetBody() *GetDatasetResponseBody
}

type GetDatasetResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponse) GoString() string {
	return s.String()
}

func (s *GetDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDatasetResponse) GetBody() *GetDatasetResponseBody {
	return s.Body
}

func (s *GetDatasetResponse) SetHeaders(v map[string]*string) *GetDatasetResponse {
	s.Headers = v
	return s
}

func (s *GetDatasetResponse) SetStatusCode(v int32) *GetDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatasetResponse) SetBody(v *GetDatasetResponseBody) *GetDatasetResponse {
	s.Body = v
	return s
}

func (s *GetDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
