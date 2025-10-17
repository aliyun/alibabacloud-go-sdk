// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterServerlessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDBClusterServerlessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDBClusterServerlessResponse
	GetStatusCode() *int32
	SetBody(v *DisableDBClusterServerlessResponseBody) *DisableDBClusterServerlessResponse
	GetBody() *DisableDBClusterServerlessResponseBody
}

type DisableDBClusterServerlessResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDBClusterServerlessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDBClusterServerlessResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterServerlessResponse) GoString() string {
	return s.String()
}

func (s *DisableDBClusterServerlessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDBClusterServerlessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDBClusterServerlessResponse) GetBody() *DisableDBClusterServerlessResponseBody {
	return s.Body
}

func (s *DisableDBClusterServerlessResponse) SetHeaders(v map[string]*string) *DisableDBClusterServerlessResponse {
	s.Headers = v
	return s
}

func (s *DisableDBClusterServerlessResponse) SetStatusCode(v int32) *DisableDBClusterServerlessResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDBClusterServerlessResponse) SetBody(v *DisableDBClusterServerlessResponseBody) *DisableDBClusterServerlessResponse {
	s.Body = v
	return s
}

func (s *DisableDBClusterServerlessResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
