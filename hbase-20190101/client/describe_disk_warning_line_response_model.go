// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskWarningLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiskWarningLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiskWarningLineResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiskWarningLineResponseBody) *DescribeDiskWarningLineResponse
	GetBody() *DescribeDiskWarningLineResponseBody
}

type DescribeDiskWarningLineResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskWarningLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskWarningLineResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskWarningLineResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskWarningLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiskWarningLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiskWarningLineResponse) GetBody() *DescribeDiskWarningLineResponseBody {
	return s.Body
}

func (s *DescribeDiskWarningLineResponse) SetHeaders(v map[string]*string) *DescribeDiskWarningLineResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskWarningLineResponse) SetStatusCode(v int32) *DescribeDiskWarningLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskWarningLineResponse) SetBody(v *DescribeDiskWarningLineResponseBody) *DescribeDiskWarningLineResponse {
	s.Body = v
	return s
}

func (s *DescribeDiskWarningLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
