// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRegistryScanDayNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetRegistryScanDayNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetRegistryScanDayNumResponse
	GetStatusCode() *int32
	SetBody(v *SetRegistryScanDayNumResponseBody) *SetRegistryScanDayNumResponse
	GetBody() *SetRegistryScanDayNumResponseBody
}

type SetRegistryScanDayNumResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRegistryScanDayNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRegistryScanDayNumResponse) String() string {
	return dara.Prettify(s)
}

func (s SetRegistryScanDayNumResponse) GoString() string {
	return s.String()
}

func (s *SetRegistryScanDayNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetRegistryScanDayNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetRegistryScanDayNumResponse) GetBody() *SetRegistryScanDayNumResponseBody {
	return s.Body
}

func (s *SetRegistryScanDayNumResponse) SetHeaders(v map[string]*string) *SetRegistryScanDayNumResponse {
	s.Headers = v
	return s
}

func (s *SetRegistryScanDayNumResponse) SetStatusCode(v int32) *SetRegistryScanDayNumResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRegistryScanDayNumResponse) SetBody(v *SetRegistryScanDayNumResponseBody) *SetRegistryScanDayNumResponse {
	s.Body = v
	return s
}

func (s *SetRegistryScanDayNumResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
