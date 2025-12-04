// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenLdpsColumnarIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenLdpsColumnarIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenLdpsColumnarIndexResponse
	GetStatusCode() *int32
	SetBody(v *OpenLdpsColumnarIndexResponseBody) *OpenLdpsColumnarIndexResponse
	GetBody() *OpenLdpsColumnarIndexResponseBody
}

type OpenLdpsColumnarIndexResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenLdpsColumnarIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenLdpsColumnarIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenLdpsColumnarIndexResponse) GoString() string {
	return s.String()
}

func (s *OpenLdpsColumnarIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenLdpsColumnarIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenLdpsColumnarIndexResponse) GetBody() *OpenLdpsColumnarIndexResponseBody {
	return s.Body
}

func (s *OpenLdpsColumnarIndexResponse) SetHeaders(v map[string]*string) *OpenLdpsColumnarIndexResponse {
	s.Headers = v
	return s
}

func (s *OpenLdpsColumnarIndexResponse) SetStatusCode(v int32) *OpenLdpsColumnarIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenLdpsColumnarIndexResponse) SetBody(v *OpenLdpsColumnarIndexResponseBody) *OpenLdpsColumnarIndexResponse {
	s.Body = v
	return s
}

func (s *OpenLdpsColumnarIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
