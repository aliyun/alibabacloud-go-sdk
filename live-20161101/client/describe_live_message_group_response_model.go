// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveMessageGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveMessageGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveMessageGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveMessageGroupResponseBody) *DescribeLiveMessageGroupResponse
	GetBody() *DescribeLiveMessageGroupResponseBody
}

type DescribeLiveMessageGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveMessageGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveMessageGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveMessageGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveMessageGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveMessageGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveMessageGroupResponse) GetBody() *DescribeLiveMessageGroupResponseBody {
	return s.Body
}

func (s *DescribeLiveMessageGroupResponse) SetHeaders(v map[string]*string) *DescribeLiveMessageGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveMessageGroupResponse) SetStatusCode(v int32) *DescribeLiveMessageGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveMessageGroupResponse) SetBody(v *DescribeLiveMessageGroupResponseBody) *DescribeLiveMessageGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveMessageGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
