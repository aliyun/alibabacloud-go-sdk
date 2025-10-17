// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublishTaskStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPublishTaskStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPublishTaskStateResponse
	GetStatusCode() *int32
	SetBody(v *GetPublishTaskStateResponseBody) *GetPublishTaskStateResponse
	GetBody() *GetPublishTaskStateResponseBody
}

type GetPublishTaskStateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublishTaskStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublishTaskStateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPublishTaskStateResponse) GoString() string {
	return s.String()
}

func (s *GetPublishTaskStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPublishTaskStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPublishTaskStateResponse) GetBody() *GetPublishTaskStateResponseBody {
	return s.Body
}

func (s *GetPublishTaskStateResponse) SetHeaders(v map[string]*string) *GetPublishTaskStateResponse {
	s.Headers = v
	return s
}

func (s *GetPublishTaskStateResponse) SetStatusCode(v int32) *GetPublishTaskStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublishTaskStateResponse) SetBody(v *GetPublishTaskStateResponseBody) *GetPublishTaskStateResponse {
	s.Body = v
	return s
}

func (s *GetPublishTaskStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
