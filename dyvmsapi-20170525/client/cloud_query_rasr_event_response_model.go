// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryRasrEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudQueryRasrEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudQueryRasrEventResponse
	GetStatusCode() *int32
	SetBody(v *CloudQueryRasrEventResponseBody) *CloudQueryRasrEventResponse
	GetBody() *CloudQueryRasrEventResponseBody
}

type CloudQueryRasrEventResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudQueryRasrEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudQueryRasrEventResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryRasrEventResponse) GoString() string {
	return s.String()
}

func (s *CloudQueryRasrEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudQueryRasrEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudQueryRasrEventResponse) GetBody() *CloudQueryRasrEventResponseBody {
	return s.Body
}

func (s *CloudQueryRasrEventResponse) SetHeaders(v map[string]*string) *CloudQueryRasrEventResponse {
	s.Headers = v
	return s
}

func (s *CloudQueryRasrEventResponse) SetStatusCode(v int32) *CloudQueryRasrEventResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryRasrEventResponse) SetBody(v *CloudQueryRasrEventResponseBody) *CloudQueryRasrEventResponse {
	s.Body = v
	return s
}

func (s *CloudQueryRasrEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
