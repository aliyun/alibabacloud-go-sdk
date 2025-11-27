// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBatchResponse
	GetStatusCode() *int32
	SetBody(v *GetBatchResponseBody) *GetBatchResponse
	GetBody() *GetBatchResponseBody
}

type GetBatchResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBatchResponse) GoString() string {
	return s.String()
}

func (s *GetBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBatchResponse) GetBody() *GetBatchResponseBody {
	return s.Body
}

func (s *GetBatchResponse) SetHeaders(v map[string]*string) *GetBatchResponse {
	s.Headers = v
	return s
}

func (s *GetBatchResponse) SetStatusCode(v int32) *GetBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBatchResponse) SetBody(v *GetBatchResponseBody) *GetBatchResponse {
	s.Body = v
	return s
}

func (s *GetBatchResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
