// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSelectionConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSelectionConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSelectionConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateSelectionConfigResponseBody) *CreateSelectionConfigResponse
	GetBody() *CreateSelectionConfigResponseBody
}

type CreateSelectionConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSelectionConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSelectionConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSelectionConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSelectionConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSelectionConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSelectionConfigResponse) GetBody() *CreateSelectionConfigResponseBody {
	return s.Body
}

func (s *CreateSelectionConfigResponse) SetHeaders(v map[string]*string) *CreateSelectionConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSelectionConfigResponse) SetStatusCode(v int32) *CreateSelectionConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSelectionConfigResponse) SetBody(v *CreateSelectionConfigResponseBody) *CreateSelectionConfigResponse {
	s.Body = v
	return s
}

func (s *CreateSelectionConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
