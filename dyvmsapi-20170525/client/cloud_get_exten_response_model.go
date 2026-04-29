// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetExtenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudGetExtenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudGetExtenResponse
	GetStatusCode() *int32
	SetBody(v *CloudGetExtenResponseBody) *CloudGetExtenResponse
	GetBody() *CloudGetExtenResponseBody
}

type CloudGetExtenResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudGetExtenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudGetExtenResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudGetExtenResponse) GoString() string {
	return s.String()
}

func (s *CloudGetExtenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudGetExtenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudGetExtenResponse) GetBody() *CloudGetExtenResponseBody {
	return s.Body
}

func (s *CloudGetExtenResponse) SetHeaders(v map[string]*string) *CloudGetExtenResponse {
	s.Headers = v
	return s
}

func (s *CloudGetExtenResponse) SetStatusCode(v int32) *CloudGetExtenResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudGetExtenResponse) SetBody(v *CloudGetExtenResponseBody) *CloudGetExtenResponse {
	s.Body = v
	return s
}

func (s *CloudGetExtenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
