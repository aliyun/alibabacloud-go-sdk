// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginLayoutDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOriginLayoutDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOriginLayoutDataResponse
	GetStatusCode() *int32
	SetBody(v *GetOriginLayoutDataResponseBody) *GetOriginLayoutDataResponse
	GetBody() *GetOriginLayoutDataResponseBody
}

type GetOriginLayoutDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOriginLayoutDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOriginLayoutDataResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOriginLayoutDataResponse) GoString() string {
	return s.String()
}

func (s *GetOriginLayoutDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOriginLayoutDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOriginLayoutDataResponse) GetBody() *GetOriginLayoutDataResponseBody {
	return s.Body
}

func (s *GetOriginLayoutDataResponse) SetHeaders(v map[string]*string) *GetOriginLayoutDataResponse {
	s.Headers = v
	return s
}

func (s *GetOriginLayoutDataResponse) SetStatusCode(v int32) *GetOriginLayoutDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOriginLayoutDataResponse) SetBody(v *GetOriginLayoutDataResponseBody) *GetOriginLayoutDataResponse {
	s.Body = v
	return s
}

func (s *GetOriginLayoutDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
