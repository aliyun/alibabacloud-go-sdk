// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventRedeployInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchEventRedeployInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchEventRedeployInstanceResponse
	GetStatusCode() *int32
	SetBody(v *BatchEventRedeployInstanceResponseBody) *BatchEventRedeployInstanceResponse
	GetBody() *BatchEventRedeployInstanceResponseBody
}

type BatchEventRedeployInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchEventRedeployInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchEventRedeployInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRedeployInstanceResponse) GoString() string {
	return s.String()
}

func (s *BatchEventRedeployInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchEventRedeployInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchEventRedeployInstanceResponse) GetBody() *BatchEventRedeployInstanceResponseBody {
	return s.Body
}

func (s *BatchEventRedeployInstanceResponse) SetHeaders(v map[string]*string) *BatchEventRedeployInstanceResponse {
	s.Headers = v
	return s
}

func (s *BatchEventRedeployInstanceResponse) SetStatusCode(v int32) *BatchEventRedeployInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchEventRedeployInstanceResponse) SetBody(v *BatchEventRedeployInstanceResponseBody) *BatchEventRedeployInstanceResponse {
	s.Body = v
	return s
}

func (s *BatchEventRedeployInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
