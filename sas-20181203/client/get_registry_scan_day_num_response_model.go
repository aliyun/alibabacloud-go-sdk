// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistryScanDayNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegistryScanDayNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegistryScanDayNumResponse
	GetStatusCode() *int32
	SetBody(v *GetRegistryScanDayNumResponseBody) *GetRegistryScanDayNumResponse
	GetBody() *GetRegistryScanDayNumResponseBody
}

type GetRegistryScanDayNumResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegistryScanDayNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegistryScanDayNumResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryScanDayNumResponse) GoString() string {
	return s.String()
}

func (s *GetRegistryScanDayNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegistryScanDayNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegistryScanDayNumResponse) GetBody() *GetRegistryScanDayNumResponseBody {
	return s.Body
}

func (s *GetRegistryScanDayNumResponse) SetHeaders(v map[string]*string) *GetRegistryScanDayNumResponse {
	s.Headers = v
	return s
}

func (s *GetRegistryScanDayNumResponse) SetStatusCode(v int32) *GetRegistryScanDayNumResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegistryScanDayNumResponse) SetBody(v *GetRegistryScanDayNumResponseBody) *GetRegistryScanDayNumResponse {
	s.Body = v
	return s
}

func (s *GetRegistryScanDayNumResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
