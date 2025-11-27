// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsumerListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsumerListResponse
	GetStatusCode() *int32
	SetBody(v *GetConsumerListResponseBody) *GetConsumerListResponse
	GetBody() *GetConsumerListResponseBody
}

type GetConsumerListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerListResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsumerListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsumerListResponse) GetBody() *GetConsumerListResponseBody {
	return s.Body
}

func (s *GetConsumerListResponse) SetHeaders(v map[string]*string) *GetConsumerListResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerListResponse) SetStatusCode(v int32) *GetConsumerListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerListResponse) SetBody(v *GetConsumerListResponseBody) *GetConsumerListResponse {
	s.Body = v
	return s
}

func (s *GetConsumerListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
