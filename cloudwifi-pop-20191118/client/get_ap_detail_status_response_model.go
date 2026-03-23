// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApDetailStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApDetailStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApDetailStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetApDetailStatusResponseBody) *GetApDetailStatusResponse
	GetBody() *GetApDetailStatusResponseBody
}

type GetApDetailStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApDetailStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApDetailStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApDetailStatusResponse) GoString() string {
	return s.String()
}

func (s *GetApDetailStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApDetailStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApDetailStatusResponse) GetBody() *GetApDetailStatusResponseBody {
	return s.Body
}

func (s *GetApDetailStatusResponse) SetHeaders(v map[string]*string) *GetApDetailStatusResponse {
	s.Headers = v
	return s
}

func (s *GetApDetailStatusResponse) SetStatusCode(v int32) *GetApDetailStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApDetailStatusResponse) SetBody(v *GetApDetailStatusResponseBody) *GetApDetailStatusResponse {
	s.Body = v
	return s
}

func (s *GetApDetailStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
