// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmartAGByAccessPointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSmartAGByAccessPointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSmartAGByAccessPointResponse
	GetStatusCode() *int32
	SetBody(v *ListSmartAGByAccessPointResponseBody) *ListSmartAGByAccessPointResponse
	GetBody() *ListSmartAGByAccessPointResponseBody
}

type ListSmartAGByAccessPointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSmartAGByAccessPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSmartAGByAccessPointResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSmartAGByAccessPointResponse) GoString() string {
	return s.String()
}

func (s *ListSmartAGByAccessPointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSmartAGByAccessPointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSmartAGByAccessPointResponse) GetBody() *ListSmartAGByAccessPointResponseBody {
	return s.Body
}

func (s *ListSmartAGByAccessPointResponse) SetHeaders(v map[string]*string) *ListSmartAGByAccessPointResponse {
	s.Headers = v
	return s
}

func (s *ListSmartAGByAccessPointResponse) SetStatusCode(v int32) *ListSmartAGByAccessPointResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSmartAGByAccessPointResponse) SetBody(v *ListSmartAGByAccessPointResponseBody) *ListSmartAGByAccessPointResponse {
	s.Body = v
	return s
}

func (s *ListSmartAGByAccessPointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
