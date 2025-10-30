// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTaskVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchTaskVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchTaskVersionsResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchTaskVersionsResponseBody) *GetBatchTaskVersionsResponse
	GetBody() *GetBatchTaskVersionsResponseBody
}

type GetBatchTaskVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchTaskVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchTaskVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTaskVersionsResponse) GoString() string {
	return s.String()
}

func (s *GetBatchTaskVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchTaskVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchTaskVersionsResponse) GetBody() *GetBatchTaskVersionsResponseBody {
	return s.Body
}

func (s *GetBatchTaskVersionsResponse) SetHeaders(v map[string]*string) *GetBatchTaskVersionsResponse {
	s.Headers = v
	return s
}

func (s *GetBatchTaskVersionsResponse) SetStatusCode(v int32) *GetBatchTaskVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchTaskVersionsResponse) SetBody(v *GetBatchTaskVersionsResponseBody) *GetBatchTaskVersionsResponse {
	s.Body = v
	return s
}

func (s *GetBatchTaskVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
