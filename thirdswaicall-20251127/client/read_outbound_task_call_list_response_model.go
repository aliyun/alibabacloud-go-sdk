// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadOutboundTaskCallListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadOutboundTaskCallListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadOutboundTaskCallListResponse
	GetStatusCode() *int32
	SetBody(v *ReadOutboundTaskCallListResponseBody) *ReadOutboundTaskCallListResponse
	GetBody() *ReadOutboundTaskCallListResponseBody
}

type ReadOutboundTaskCallListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadOutboundTaskCallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadOutboundTaskCallListResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadOutboundTaskCallListResponse) GoString() string {
	return s.String()
}

func (s *ReadOutboundTaskCallListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadOutboundTaskCallListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadOutboundTaskCallListResponse) GetBody() *ReadOutboundTaskCallListResponseBody {
	return s.Body
}

func (s *ReadOutboundTaskCallListResponse) SetHeaders(v map[string]*string) *ReadOutboundTaskCallListResponse {
	s.Headers = v
	return s
}

func (s *ReadOutboundTaskCallListResponse) SetStatusCode(v int32) *ReadOutboundTaskCallListResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadOutboundTaskCallListResponse) SetBody(v *ReadOutboundTaskCallListResponseBody) *ReadOutboundTaskCallListResponse {
	s.Body = v
	return s
}

func (s *ReadOutboundTaskCallListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
