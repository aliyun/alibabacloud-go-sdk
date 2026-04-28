// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudQueryWebcallCdrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudQueryWebcallCdrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudQueryWebcallCdrResponse
	GetStatusCode() *int32
	SetBody(v *CloudQueryWebcallCdrResponseBody) *CloudQueryWebcallCdrResponse
	GetBody() *CloudQueryWebcallCdrResponseBody
}

type CloudQueryWebcallCdrResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudQueryWebcallCdrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudQueryWebcallCdrResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudQueryWebcallCdrResponse) GoString() string {
	return s.String()
}

func (s *CloudQueryWebcallCdrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudQueryWebcallCdrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudQueryWebcallCdrResponse) GetBody() *CloudQueryWebcallCdrResponseBody {
	return s.Body
}

func (s *CloudQueryWebcallCdrResponse) SetHeaders(v map[string]*string) *CloudQueryWebcallCdrResponse {
	s.Headers = v
	return s
}

func (s *CloudQueryWebcallCdrResponse) SetStatusCode(v int32) *CloudQueryWebcallCdrResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudQueryWebcallCdrResponse) SetBody(v *CloudQueryWebcallCdrResponseBody) *CloudQueryWebcallCdrResponse {
	s.Body = v
	return s
}

func (s *CloudQueryWebcallCdrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
