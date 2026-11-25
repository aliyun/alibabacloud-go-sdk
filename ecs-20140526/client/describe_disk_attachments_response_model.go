// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiskAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiskAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiskAttachmentsResponseBody) *DescribeDiskAttachmentsResponse
	GetBody() *DescribeDiskAttachmentsResponseBody
}

type DescribeDiskAttachmentsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiskAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiskAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiskAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiskAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiskAttachmentsResponse) GetBody() *DescribeDiskAttachmentsResponseBody {
	return s.Body
}

func (s *DescribeDiskAttachmentsResponse) SetHeaders(v map[string]*string) *DescribeDiskAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiskAttachmentsResponse) SetStatusCode(v int32) *DescribeDiskAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiskAttachmentsResponse) SetBody(v *DescribeDiskAttachmentsResponseBody) *DescribeDiskAttachmentsResponse {
	s.Body = v
	return s
}

func (s *DescribeDiskAttachmentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
