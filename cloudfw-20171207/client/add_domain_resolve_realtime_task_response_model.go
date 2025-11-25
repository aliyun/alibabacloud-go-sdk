// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainResolveRealtimeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDomainResolveRealtimeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDomainResolveRealtimeTaskResponse
	GetStatusCode() *int32
	SetBody(v *AddDomainResolveRealtimeTaskResponseBody) *AddDomainResolveRealtimeTaskResponse
	GetBody() *AddDomainResolveRealtimeTaskResponseBody
}

type AddDomainResolveRealtimeTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDomainResolveRealtimeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDomainResolveRealtimeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDomainResolveRealtimeTaskResponse) GoString() string {
	return s.String()
}

func (s *AddDomainResolveRealtimeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDomainResolveRealtimeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDomainResolveRealtimeTaskResponse) GetBody() *AddDomainResolveRealtimeTaskResponseBody {
	return s.Body
}

func (s *AddDomainResolveRealtimeTaskResponse) SetHeaders(v map[string]*string) *AddDomainResolveRealtimeTaskResponse {
	s.Headers = v
	return s
}

func (s *AddDomainResolveRealtimeTaskResponse) SetStatusCode(v int32) *AddDomainResolveRealtimeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDomainResolveRealtimeTaskResponse) SetBody(v *AddDomainResolveRealtimeTaskResponseBody) *AddDomainResolveRealtimeTaskResponse {
	s.Body = v
	return s
}

func (s *AddDomainResolveRealtimeTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
