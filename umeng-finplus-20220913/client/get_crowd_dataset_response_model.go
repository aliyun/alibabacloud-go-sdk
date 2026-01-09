// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCrowdDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCrowdDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCrowdDatasetResponse
	GetStatusCode() *int32
	SetBody(v *GetCrowdDatasetResponseBody) *GetCrowdDatasetResponse
	GetBody() *GetCrowdDatasetResponseBody
}

type GetCrowdDatasetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCrowdDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCrowdDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCrowdDatasetResponse) GoString() string {
	return s.String()
}

func (s *GetCrowdDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCrowdDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCrowdDatasetResponse) GetBody() *GetCrowdDatasetResponseBody {
	return s.Body
}

func (s *GetCrowdDatasetResponse) SetHeaders(v map[string]*string) *GetCrowdDatasetResponse {
	s.Headers = v
	return s
}

func (s *GetCrowdDatasetResponse) SetStatusCode(v int32) *GetCrowdDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCrowdDatasetResponse) SetBody(v *GetCrowdDatasetResponseBody) *GetCrowdDatasetResponse {
	s.Body = v
	return s
}

func (s *GetCrowdDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
