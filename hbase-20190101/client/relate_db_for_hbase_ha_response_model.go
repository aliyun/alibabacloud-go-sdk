// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRelateDbForHBaseHaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RelateDbForHBaseHaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RelateDbForHBaseHaResponse
	GetStatusCode() *int32
	SetBody(v *RelateDbForHBaseHaResponseBody) *RelateDbForHBaseHaResponse
	GetBody() *RelateDbForHBaseHaResponseBody
}

type RelateDbForHBaseHaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RelateDbForHBaseHaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RelateDbForHBaseHaResponse) String() string {
	return dara.Prettify(s)
}

func (s RelateDbForHBaseHaResponse) GoString() string {
	return s.String()
}

func (s *RelateDbForHBaseHaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RelateDbForHBaseHaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RelateDbForHBaseHaResponse) GetBody() *RelateDbForHBaseHaResponseBody {
	return s.Body
}

func (s *RelateDbForHBaseHaResponse) SetHeaders(v map[string]*string) *RelateDbForHBaseHaResponse {
	s.Headers = v
	return s
}

func (s *RelateDbForHBaseHaResponse) SetStatusCode(v int32) *RelateDbForHBaseHaResponse {
	s.StatusCode = &v
	return s
}

func (s *RelateDbForHBaseHaResponse) SetBody(v *RelateDbForHBaseHaResponseBody) *RelateDbForHBaseHaResponse {
	s.Body = v
	return s
}

func (s *RelateDbForHBaseHaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
