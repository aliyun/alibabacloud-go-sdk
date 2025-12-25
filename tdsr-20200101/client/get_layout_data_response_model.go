// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLayoutDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLayoutDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLayoutDataResponse
	GetStatusCode() *int32
	SetBody(v *GetLayoutDataResponseBody) *GetLayoutDataResponse
	GetBody() *GetLayoutDataResponseBody
}

type GetLayoutDataResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLayoutDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLayoutDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLayoutDataResponse) GoString() string {
	return s.String()
}

func (s *GetLayoutDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLayoutDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLayoutDataResponse) GetBody() *GetLayoutDataResponseBody {
	return s.Body
}

func (s *GetLayoutDataResponse) SetHeaders(v map[string]*string) *GetLayoutDataResponse {
	s.Headers = v
	return s
}

func (s *GetLayoutDataResponse) SetStatusCode(v int32) *GetLayoutDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLayoutDataResponse) SetBody(v *GetLayoutDataResponseBody) *GetLayoutDataResponse {
	s.Body = v
	return s
}

func (s *GetLayoutDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
