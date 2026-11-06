// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterCTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDBClusterCTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDBClusterCTableResponse
	GetStatusCode() *int32
	SetBody(v *DisableDBClusterCTableResponseBody) *DisableDBClusterCTableResponse
	GetBody() *DisableDBClusterCTableResponseBody
}

type DisableDBClusterCTableResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDBClusterCTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDBClusterCTableResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterCTableResponse) GoString() string {
	return s.String()
}

func (s *DisableDBClusterCTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDBClusterCTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDBClusterCTableResponse) GetBody() *DisableDBClusterCTableResponseBody {
	return s.Body
}

func (s *DisableDBClusterCTableResponse) SetHeaders(v map[string]*string) *DisableDBClusterCTableResponse {
	s.Headers = v
	return s
}

func (s *DisableDBClusterCTableResponse) SetStatusCode(v int32) *DisableDBClusterCTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDBClusterCTableResponse) SetBody(v *DisableDBClusterCTableResponseBody) *DisableDBClusterCTableResponse {
	s.Body = v
	return s
}

func (s *DisableDBClusterCTableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
