// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSdlLastPayloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSdlLastPayloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSdlLastPayloadResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSdlLastPayloadResponseBody) *DescribeSdlLastPayloadResponse
	GetBody() *DescribeSdlLastPayloadResponseBody
}

type DescribeSdlLastPayloadResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSdlLastPayloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSdlLastPayloadResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSdlLastPayloadResponse) GoString() string {
	return s.String()
}

func (s *DescribeSdlLastPayloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSdlLastPayloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSdlLastPayloadResponse) GetBody() *DescribeSdlLastPayloadResponseBody {
	return s.Body
}

func (s *DescribeSdlLastPayloadResponse) SetHeaders(v map[string]*string) *DescribeSdlLastPayloadResponse {
	s.Headers = v
	return s
}

func (s *DescribeSdlLastPayloadResponse) SetStatusCode(v int32) *DescribeSdlLastPayloadResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSdlLastPayloadResponse) SetBody(v *DescribeSdlLastPayloadResponseBody) *DescribeSdlLastPayloadResponse {
	s.Body = v
	return s
}

func (s *DescribeSdlLastPayloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
