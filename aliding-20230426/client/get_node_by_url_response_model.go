// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeByUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeByUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeByUrlResponseBody) *GetNodeByUrlResponse
	GetBody() *GetNodeByUrlResponseBody
}

type GetNodeByUrlResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeByUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeByUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByUrlResponse) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeByUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeByUrlResponse) GetBody() *GetNodeByUrlResponseBody {
	return s.Body
}

func (s *GetNodeByUrlResponse) SetHeaders(v map[string]*string) *GetNodeByUrlResponse {
	s.Headers = v
	return s
}

func (s *GetNodeByUrlResponse) SetStatusCode(v int32) *GetNodeByUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeByUrlResponse) SetBody(v *GetNodeByUrlResponseBody) *GetNodeByUrlResponse {
	s.Body = v
	return s
}

func (s *GetNodeByUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
