// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventHouseRuntimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEventHouseRuntimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEventHouseRuntimeResponse
	GetStatusCode() *int32
	SetBody(v *GetEventHouseRuntimeResponseBody) *GetEventHouseRuntimeResponse
	GetBody() *GetEventHouseRuntimeResponseBody
}

type GetEventHouseRuntimeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventHouseRuntimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventHouseRuntimeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEventHouseRuntimeResponse) GoString() string {
	return s.String()
}

func (s *GetEventHouseRuntimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEventHouseRuntimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEventHouseRuntimeResponse) GetBody() *GetEventHouseRuntimeResponseBody {
	return s.Body
}

func (s *GetEventHouseRuntimeResponse) SetHeaders(v map[string]*string) *GetEventHouseRuntimeResponse {
	s.Headers = v
	return s
}

func (s *GetEventHouseRuntimeResponse) SetStatusCode(v int32) *GetEventHouseRuntimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventHouseRuntimeResponse) SetBody(v *GetEventHouseRuntimeResponseBody) *GetEventHouseRuntimeResponse {
	s.Body = v
	return s
}

func (s *GetEventHouseRuntimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
