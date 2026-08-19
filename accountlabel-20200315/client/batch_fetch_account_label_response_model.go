// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchFetchAccountLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchFetchAccountLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchFetchAccountLabelResponse
	GetStatusCode() *int32
	SetBody(v *BatchFetchAccountLabelResponseBody) *BatchFetchAccountLabelResponse
	GetBody() *BatchFetchAccountLabelResponseBody
}

type BatchFetchAccountLabelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchFetchAccountLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchFetchAccountLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchFetchAccountLabelResponse) GoString() string {
	return s.String()
}

func (s *BatchFetchAccountLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchFetchAccountLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchFetchAccountLabelResponse) GetBody() *BatchFetchAccountLabelResponseBody {
	return s.Body
}

func (s *BatchFetchAccountLabelResponse) SetHeaders(v map[string]*string) *BatchFetchAccountLabelResponse {
	s.Headers = v
	return s
}

func (s *BatchFetchAccountLabelResponse) SetStatusCode(v int32) *BatchFetchAccountLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchFetchAccountLabelResponse) SetBody(v *BatchFetchAccountLabelResponseBody) *BatchFetchAccountLabelResponse {
	s.Body = v
	return s
}

func (s *BatchFetchAccountLabelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
