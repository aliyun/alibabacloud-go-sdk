// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryObCdrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudQueryObCdrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudQueryObCdrResponse
	GetStatusCode() *int32
	SetBody(v *CloudQueryObCdrResponseBody) *CloudQueryObCdrResponse
	GetBody() *CloudQueryObCdrResponseBody
}

type CloudQueryObCdrResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudQueryObCdrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudQueryObCdrResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryObCdrResponse) GoString() string {
	return s.String()
}

func (s *CloudQueryObCdrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudQueryObCdrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudQueryObCdrResponse) GetBody() *CloudQueryObCdrResponseBody {
	return s.Body
}

func (s *CloudQueryObCdrResponse) SetHeaders(v map[string]*string) *CloudQueryObCdrResponse {
	s.Headers = v
	return s
}

func (s *CloudQueryObCdrResponse) SetStatusCode(v int32) *CloudQueryObCdrResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryObCdrResponse) SetBody(v *CloudQueryObCdrResponseBody) *CloudQueryObCdrResponse {
	s.Body = v
	return s
}

func (s *CloudQueryObCdrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
