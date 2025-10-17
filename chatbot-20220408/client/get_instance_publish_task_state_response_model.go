// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstancePublishTaskStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstancePublishTaskStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstancePublishTaskStateResponse
	GetStatusCode() *int32
	SetBody(v *GetInstancePublishTaskStateResponseBody) *GetInstancePublishTaskStateResponse
	GetBody() *GetInstancePublishTaskStateResponseBody
}

type GetInstancePublishTaskStateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstancePublishTaskStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstancePublishTaskStateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstancePublishTaskStateResponse) GoString() string {
	return s.String()
}

func (s *GetInstancePublishTaskStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstancePublishTaskStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstancePublishTaskStateResponse) GetBody() *GetInstancePublishTaskStateResponseBody {
	return s.Body
}

func (s *GetInstancePublishTaskStateResponse) SetHeaders(v map[string]*string) *GetInstancePublishTaskStateResponse {
	s.Headers = v
	return s
}

func (s *GetInstancePublishTaskStateResponse) SetStatusCode(v int32) *GetInstancePublishTaskStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstancePublishTaskStateResponse) SetBody(v *GetInstancePublishTaskStateResponseBody) *GetInstancePublishTaskStateResponse {
	s.Body = v
	return s
}

func (s *GetInstancePublishTaskStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
