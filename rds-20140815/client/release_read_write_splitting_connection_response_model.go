// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseReadWriteSplittingConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseReadWriteSplittingConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseReadWriteSplittingConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseReadWriteSplittingConnectionResponseBody) *ReleaseReadWriteSplittingConnectionResponse
	GetBody() *ReleaseReadWriteSplittingConnectionResponseBody
}

type ReleaseReadWriteSplittingConnectionResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseReadWriteSplittingConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseReadWriteSplittingConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseReadWriteSplittingConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseReadWriteSplittingConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseReadWriteSplittingConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseReadWriteSplittingConnectionResponse) GetBody() *ReleaseReadWriteSplittingConnectionResponseBody {
	return s.Body
}

func (s *ReleaseReadWriteSplittingConnectionResponse) SetHeaders(v map[string]*string) *ReleaseReadWriteSplittingConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseReadWriteSplittingConnectionResponse) SetStatusCode(v int32) *ReleaseReadWriteSplittingConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseReadWriteSplittingConnectionResponse) SetBody(v *ReleaseReadWriteSplittingConnectionResponseBody) *ReleaseReadWriteSplittingConnectionResponse {
	s.Body = v
	return s
}

func (s *ReleaseReadWriteSplittingConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
