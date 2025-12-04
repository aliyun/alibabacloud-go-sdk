// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDataEventSelectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutDataEventSelectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutDataEventSelectorResponse
	GetStatusCode() *int32
	SetBody(v *PutDataEventSelectorResponseBody) *PutDataEventSelectorResponse
	GetBody() *PutDataEventSelectorResponseBody
}

type PutDataEventSelectorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutDataEventSelectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutDataEventSelectorResponse) String() string {
	return dara.Prettify(s)
}

func (s PutDataEventSelectorResponse) GoString() string {
	return s.String()
}

func (s *PutDataEventSelectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutDataEventSelectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutDataEventSelectorResponse) GetBody() *PutDataEventSelectorResponseBody {
	return s.Body
}

func (s *PutDataEventSelectorResponse) SetHeaders(v map[string]*string) *PutDataEventSelectorResponse {
	s.Headers = v
	return s
}

func (s *PutDataEventSelectorResponse) SetStatusCode(v int32) *PutDataEventSelectorResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDataEventSelectorResponse) SetBody(v *PutDataEventSelectorResponseBody) *PutDataEventSelectorResponse {
	s.Body = v
	return s
}

func (s *PutDataEventSelectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
