// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListIvrsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClinkListIvrsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClinkListIvrsResponse
	GetStatusCode() *int32
	SetBody(v *ClinkListIvrsResponseBody) *ClinkListIvrsResponse
	GetBody() *ClinkListIvrsResponseBody
}

type ClinkListIvrsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClinkListIvrsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClinkListIvrsResponse) String() string {
	return dara.Prettify(s)
}

func (s ClinkListIvrsResponse) GoString() string {
	return s.String()
}

func (s *ClinkListIvrsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClinkListIvrsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClinkListIvrsResponse) GetBody() *ClinkListIvrsResponseBody {
	return s.Body
}

func (s *ClinkListIvrsResponse) SetHeaders(v map[string]*string) *ClinkListIvrsResponse {
	s.Headers = v
	return s
}

func (s *ClinkListIvrsResponse) SetStatusCode(v int32) *ClinkListIvrsResponse {
	s.StatusCode = &v
	return s
}

func (s *ClinkListIvrsResponse) SetBody(v *ClinkListIvrsResponseBody) *ClinkListIvrsResponse {
	s.Body = v
	return s
}

func (s *ClinkListIvrsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
