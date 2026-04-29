// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListExtenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListExtenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListExtenResponse
	GetStatusCode() *int32
	SetBody(v *CloudListExtenResponseBody) *CloudListExtenResponse
	GetBody() *CloudListExtenResponseBody
}

type CloudListExtenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListExtenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListExtenResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListExtenResponse) GoString() string {
	return s.String()
}

func (s *CloudListExtenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListExtenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListExtenResponse) GetBody() *CloudListExtenResponseBody {
	return s.Body
}

func (s *CloudListExtenResponse) SetHeaders(v map[string]*string) *CloudListExtenResponse {
	s.Headers = v
	return s
}

func (s *CloudListExtenResponse) SetStatusCode(v int32) *CloudListExtenResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListExtenResponse) SetBody(v *CloudListExtenResponseBody) *CloudListExtenResponse {
	s.Body = v
	return s
}

func (s *CloudListExtenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
