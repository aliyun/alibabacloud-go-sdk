// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceControlEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceControlEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceControlEventResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceControlEventResponseBody) *GetResourceControlEventResponse
	GetBody() *GetResourceControlEventResponseBody
}

type GetResourceControlEventResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceControlEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceControlEventResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceControlEventResponse) GoString() string {
	return s.String()
}

func (s *GetResourceControlEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceControlEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceControlEventResponse) GetBody() *GetResourceControlEventResponseBody {
	return s.Body
}

func (s *GetResourceControlEventResponse) SetHeaders(v map[string]*string) *GetResourceControlEventResponse {
	s.Headers = v
	return s
}

func (s *GetResourceControlEventResponse) SetStatusCode(v int32) *GetResourceControlEventResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceControlEventResponse) SetBody(v *GetResourceControlEventResponseBody) *GetResourceControlEventResponse {
	s.Body = v
	return s
}

func (s *GetResourceControlEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
