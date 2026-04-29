// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudGetAsrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudGetAsrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudGetAsrResponse
	GetStatusCode() *int32
	SetBody(v *CloudGetAsrResponseBody) *CloudGetAsrResponse
	GetBody() *CloudGetAsrResponseBody
}

type CloudGetAsrResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudGetAsrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudGetAsrResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudGetAsrResponse) GoString() string {
	return s.String()
}

func (s *CloudGetAsrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudGetAsrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudGetAsrResponse) GetBody() *CloudGetAsrResponseBody {
	return s.Body
}

func (s *CloudGetAsrResponse) SetHeaders(v map[string]*string) *CloudGetAsrResponse {
	s.Headers = v
	return s
}

func (s *CloudGetAsrResponse) SetStatusCode(v int32) *CloudGetAsrResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudGetAsrResponse) SetBody(v *CloudGetAsrResponseBody) *CloudGetAsrResponse {
	s.Body = v
	return s
}

func (s *CloudGetAsrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
