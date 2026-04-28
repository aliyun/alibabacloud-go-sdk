// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetObCdrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudGetObCdrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudGetObCdrResponse
	GetStatusCode() *int32
	SetBody(v *CloudGetObCdrResponseBody) *CloudGetObCdrResponse
	GetBody() *CloudGetObCdrResponseBody
}

type CloudGetObCdrResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudGetObCdrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudGetObCdrResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudGetObCdrResponse) GoString() string {
	return s.String()
}

func (s *CloudGetObCdrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudGetObCdrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudGetObCdrResponse) GetBody() *CloudGetObCdrResponseBody {
	return s.Body
}

func (s *CloudGetObCdrResponse) SetHeaders(v map[string]*string) *CloudGetObCdrResponse {
	s.Headers = v
	return s
}

func (s *CloudGetObCdrResponse) SetStatusCode(v int32) *CloudGetObCdrResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudGetObCdrResponse) SetBody(v *CloudGetObCdrResponseBody) *CloudGetObCdrResponse {
	s.Body = v
	return s
}

func (s *CloudGetObCdrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
