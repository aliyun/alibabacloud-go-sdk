// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRemoteAccessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSagRemoteAccessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSagRemoteAccessResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSagRemoteAccessResponseBody) *DescribeSagRemoteAccessResponse
	GetBody() *DescribeSagRemoteAccessResponseBody
}

type DescribeSagRemoteAccessResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSagRemoteAccessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSagRemoteAccessResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRemoteAccessResponse) GoString() string {
	return s.String()
}

func (s *DescribeSagRemoteAccessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSagRemoteAccessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSagRemoteAccessResponse) GetBody() *DescribeSagRemoteAccessResponseBody {
	return s.Body
}

func (s *DescribeSagRemoteAccessResponse) SetHeaders(v map[string]*string) *DescribeSagRemoteAccessResponse {
	s.Headers = v
	return s
}

func (s *DescribeSagRemoteAccessResponse) SetStatusCode(v int32) *DescribeSagRemoteAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSagRemoteAccessResponse) SetBody(v *DescribeSagRemoteAccessResponseBody) *DescribeSagRemoteAccessResponse {
	s.Body = v
	return s
}

func (s *DescribeSagRemoteAccessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
