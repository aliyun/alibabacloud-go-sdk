// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCronClearConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataCronClearConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataCronClearConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetDataCronClearConfigResponseBody) *GetDataCronClearConfigResponse
	GetBody() *GetDataCronClearConfigResponseBody
}

type GetDataCronClearConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataCronClearConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataCronClearConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataCronClearConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDataCronClearConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataCronClearConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataCronClearConfigResponse) GetBody() *GetDataCronClearConfigResponseBody {
	return s.Body
}

func (s *GetDataCronClearConfigResponse) SetHeaders(v map[string]*string) *GetDataCronClearConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDataCronClearConfigResponse) SetStatusCode(v int32) *GetDataCronClearConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataCronClearConfigResponse) SetBody(v *GetDataCronClearConfigResponseBody) *GetDataCronClearConfigResponse {
	s.Body = v
	return s
}

func (s *GetDataCronClearConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
