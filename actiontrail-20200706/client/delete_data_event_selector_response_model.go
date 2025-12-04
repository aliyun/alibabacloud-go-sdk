// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataEventSelectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDataEventSelectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDataEventSelectorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDataEventSelectorResponseBody) *DeleteDataEventSelectorResponse
	GetBody() *DeleteDataEventSelectorResponseBody
}

type DeleteDataEventSelectorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDataEventSelectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDataEventSelectorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataEventSelectorResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataEventSelectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDataEventSelectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDataEventSelectorResponse) GetBody() *DeleteDataEventSelectorResponseBody {
	return s.Body
}

func (s *DeleteDataEventSelectorResponse) SetHeaders(v map[string]*string) *DeleteDataEventSelectorResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataEventSelectorResponse) SetStatusCode(v int32) *DeleteDataEventSelectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataEventSelectorResponse) SetBody(v *DeleteDataEventSelectorResponseBody) *DeleteDataEventSelectorResponse {
	s.Body = v
	return s
}

func (s *DeleteDataEventSelectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
