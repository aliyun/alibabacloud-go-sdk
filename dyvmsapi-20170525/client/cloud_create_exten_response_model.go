// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateExtenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudCreateExtenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudCreateExtenResponse
	GetStatusCode() *int32
	SetBody(v *CloudCreateExtenResponseBody) *CloudCreateExtenResponse
	GetBody() *CloudCreateExtenResponseBody
}

type CloudCreateExtenResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudCreateExtenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudCreateExtenResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateExtenResponse) GoString() string {
	return s.String()
}

func (s *CloudCreateExtenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudCreateExtenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudCreateExtenResponse) GetBody() *CloudCreateExtenResponseBody {
	return s.Body
}

func (s *CloudCreateExtenResponse) SetHeaders(v map[string]*string) *CloudCreateExtenResponse {
	s.Headers = v
	return s
}

func (s *CloudCreateExtenResponse) SetStatusCode(v int32) *CloudCreateExtenResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudCreateExtenResponse) SetBody(v *CloudCreateExtenResponseBody) *CloudCreateExtenResponse {
	s.Body = v
	return s
}

func (s *CloudCreateExtenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
