// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryIndividuationTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchQueryIndividuationTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchQueryIndividuationTextResponse
	GetStatusCode() *int32
	SetBody(v *BatchQueryIndividuationTextResponseBody) *BatchQueryIndividuationTextResponse
	GetBody() *BatchQueryIndividuationTextResponseBody
}

type BatchQueryIndividuationTextResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryIndividuationTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryIndividuationTextResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryIndividuationTextResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryIndividuationTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchQueryIndividuationTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchQueryIndividuationTextResponse) GetBody() *BatchQueryIndividuationTextResponseBody {
	return s.Body
}

func (s *BatchQueryIndividuationTextResponse) SetHeaders(v map[string]*string) *BatchQueryIndividuationTextResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryIndividuationTextResponse) SetStatusCode(v int32) *BatchQueryIndividuationTextResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryIndividuationTextResponse) SetBody(v *BatchQueryIndividuationTextResponseBody) *BatchQueryIndividuationTextResponse {
	s.Body = v
	return s
}

func (s *BatchQueryIndividuationTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
