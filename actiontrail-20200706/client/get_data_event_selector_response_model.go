// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataEventSelectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataEventSelectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataEventSelectorResponse
	GetStatusCode() *int32
	SetBody(v *GetDataEventSelectorResponseBody) *GetDataEventSelectorResponse
	GetBody() *GetDataEventSelectorResponseBody
}

type GetDataEventSelectorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataEventSelectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataEventSelectorResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataEventSelectorResponse) GoString() string {
	return s.String()
}

func (s *GetDataEventSelectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataEventSelectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataEventSelectorResponse) GetBody() *GetDataEventSelectorResponseBody {
	return s.Body
}

func (s *GetDataEventSelectorResponse) SetHeaders(v map[string]*string) *GetDataEventSelectorResponse {
	s.Headers = v
	return s
}

func (s *GetDataEventSelectorResponse) SetStatusCode(v int32) *GetDataEventSelectorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataEventSelectorResponse) SetBody(v *GetDataEventSelectorResponseBody) *GetDataEventSelectorResponse {
	s.Body = v
	return s
}

func (s *GetDataEventSelectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
