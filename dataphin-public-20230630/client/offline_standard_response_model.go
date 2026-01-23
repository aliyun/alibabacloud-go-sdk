// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineStandardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OfflineStandardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OfflineStandardResponse
	GetStatusCode() *int32
	SetBody(v *OfflineStandardResponseBody) *OfflineStandardResponse
	GetBody() *OfflineStandardResponseBody
}

type OfflineStandardResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OfflineStandardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OfflineStandardResponse) String() string {
	return dara.Prettify(s)
}

func (s OfflineStandardResponse) GoString() string {
	return s.String()
}

func (s *OfflineStandardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OfflineStandardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OfflineStandardResponse) GetBody() *OfflineStandardResponseBody {
	return s.Body
}

func (s *OfflineStandardResponse) SetHeaders(v map[string]*string) *OfflineStandardResponse {
	s.Headers = v
	return s
}

func (s *OfflineStandardResponse) SetStatusCode(v int32) *OfflineStandardResponse {
	s.StatusCode = &v
	return s
}

func (s *OfflineStandardResponse) SetBody(v *OfflineStandardResponseBody) *OfflineStandardResponse {
	s.Body = v
	return s
}

func (s *OfflineStandardResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
