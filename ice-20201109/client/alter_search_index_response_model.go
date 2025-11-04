// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterSearchIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AlterSearchIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AlterSearchIndexResponse
	GetStatusCode() *int32
	SetBody(v *AlterSearchIndexResponseBody) *AlterSearchIndexResponse
	GetBody() *AlterSearchIndexResponseBody
}

type AlterSearchIndexResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AlterSearchIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AlterSearchIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s AlterSearchIndexResponse) GoString() string {
	return s.String()
}

func (s *AlterSearchIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AlterSearchIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AlterSearchIndexResponse) GetBody() *AlterSearchIndexResponseBody {
	return s.Body
}

func (s *AlterSearchIndexResponse) SetHeaders(v map[string]*string) *AlterSearchIndexResponse {
	s.Headers = v
	return s
}

func (s *AlterSearchIndexResponse) SetStatusCode(v int32) *AlterSearchIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *AlterSearchIndexResponse) SetBody(v *AlterSearchIndexResponseBody) *AlterSearchIndexResponse {
	s.Body = v
	return s
}

func (s *AlterSearchIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
