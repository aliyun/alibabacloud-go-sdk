// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateAsrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudCreateAsrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudCreateAsrResponse
	GetStatusCode() *int32
	SetBody(v *CloudCreateAsrResponseBody) *CloudCreateAsrResponse
	GetBody() *CloudCreateAsrResponseBody
}

type CloudCreateAsrResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudCreateAsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudCreateAsrResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateAsrResponse) GoString() string {
	return s.String()
}

func (s *CloudCreateAsrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudCreateAsrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudCreateAsrResponse) GetBody() *CloudCreateAsrResponseBody {
	return s.Body
}

func (s *CloudCreateAsrResponse) SetHeaders(v map[string]*string) *CloudCreateAsrResponse {
	s.Headers = v
	return s
}

func (s *CloudCreateAsrResponse) SetStatusCode(v int32) *CloudCreateAsrResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudCreateAsrResponse) SetBody(v *CloudCreateAsrResponseBody) *CloudCreateAsrResponse {
	s.Body = v
	return s
}

func (s *CloudCreateAsrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
