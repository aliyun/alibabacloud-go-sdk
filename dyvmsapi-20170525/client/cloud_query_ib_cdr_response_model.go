// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryIbCdrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudQueryIbCdrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudQueryIbCdrResponse
	GetStatusCode() *int32
	SetBody(v *CloudQueryIbCdrResponseBody) *CloudQueryIbCdrResponse
	GetBody() *CloudQueryIbCdrResponseBody
}

type CloudQueryIbCdrResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudQueryIbCdrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudQueryIbCdrResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryIbCdrResponse) GoString() string {
	return s.String()
}

func (s *CloudQueryIbCdrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudQueryIbCdrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudQueryIbCdrResponse) GetBody() *CloudQueryIbCdrResponseBody {
	return s.Body
}

func (s *CloudQueryIbCdrResponse) SetHeaders(v map[string]*string) *CloudQueryIbCdrResponse {
	s.Headers = v
	return s
}

func (s *CloudQueryIbCdrResponse) SetStatusCode(v int32) *CloudQueryIbCdrResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryIbCdrResponse) SetBody(v *CloudQueryIbCdrResponseBody) *CloudQueryIbCdrResponse {
	s.Body = v
	return s
}

func (s *CloudQueryIbCdrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
