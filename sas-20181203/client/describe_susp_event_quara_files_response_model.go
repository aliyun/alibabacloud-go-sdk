// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspEventQuaraFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSuspEventQuaraFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSuspEventQuaraFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSuspEventQuaraFilesResponseBody) *DescribeSuspEventQuaraFilesResponse
	GetBody() *DescribeSuspEventQuaraFilesResponseBody
}

type DescribeSuspEventQuaraFilesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSuspEventQuaraFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSuspEventQuaraFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspEventQuaraFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspEventQuaraFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSuspEventQuaraFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSuspEventQuaraFilesResponse) GetBody() *DescribeSuspEventQuaraFilesResponseBody {
	return s.Body
}

func (s *DescribeSuspEventQuaraFilesResponse) SetHeaders(v map[string]*string) *DescribeSuspEventQuaraFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponse) SetStatusCode(v int32) *DescribeSuspEventQuaraFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponse) SetBody(v *DescribeSuspEventQuaraFilesResponseBody) *DescribeSuspEventQuaraFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeSuspEventQuaraFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
