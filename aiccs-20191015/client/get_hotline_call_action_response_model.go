// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineCallActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHotlineCallActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHotlineCallActionResponse
	GetStatusCode() *int32
	SetBody(v *GetHotlineCallActionResponseBody) *GetHotlineCallActionResponse
	GetBody() *GetHotlineCallActionResponseBody
}

type GetHotlineCallActionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHotlineCallActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHotlineCallActionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineCallActionResponse) GoString() string {
	return s.String()
}

func (s *GetHotlineCallActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHotlineCallActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotlineCallActionResponse) GetBody() *GetHotlineCallActionResponseBody {
	return s.Body
}

func (s *GetHotlineCallActionResponse) SetHeaders(v map[string]*string) *GetHotlineCallActionResponse {
	s.Headers = v
	return s
}

func (s *GetHotlineCallActionResponse) SetStatusCode(v int32) *GetHotlineCallActionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHotlineCallActionResponse) SetBody(v *GetHotlineCallActionResponseBody) *GetHotlineCallActionResponse {
	s.Body = v
	return s
}

func (s *GetHotlineCallActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
