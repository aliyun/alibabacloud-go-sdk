// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneServiceResponse
	GetStatusCode() *int32
	SetBody(v *CloneServiceResponseBody) *CloneServiceResponse
	GetBody() *CloneServiceResponseBody
}

type CloneServiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneServiceResponse) GoString() string {
	return s.String()
}

func (s *CloneServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneServiceResponse) GetBody() *CloneServiceResponseBody {
	return s.Body
}

func (s *CloneServiceResponse) SetHeaders(v map[string]*string) *CloneServiceResponse {
	s.Headers = v
	return s
}

func (s *CloneServiceResponse) SetStatusCode(v int32) *CloneServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneServiceResponse) SetBody(v *CloneServiceResponseBody) *CloneServiceResponse {
	s.Body = v
	return s
}

func (s *CloneServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
