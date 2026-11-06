// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterDocumentDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDBClusterDocumentDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDBClusterDocumentDBResponse
	GetStatusCode() *int32
	SetBody(v *DisableDBClusterDocumentDBResponseBody) *DisableDBClusterDocumentDBResponse
	GetBody() *DisableDBClusterDocumentDBResponseBody
}

type DisableDBClusterDocumentDBResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDBClusterDocumentDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDBClusterDocumentDBResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterDocumentDBResponse) GoString() string {
	return s.String()
}

func (s *DisableDBClusterDocumentDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDBClusterDocumentDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDBClusterDocumentDBResponse) GetBody() *DisableDBClusterDocumentDBResponseBody {
	return s.Body
}

func (s *DisableDBClusterDocumentDBResponse) SetHeaders(v map[string]*string) *DisableDBClusterDocumentDBResponse {
	s.Headers = v
	return s
}

func (s *DisableDBClusterDocumentDBResponse) SetStatusCode(v int32) *DisableDBClusterDocumentDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDBClusterDocumentDBResponse) SetBody(v *DisableDBClusterDocumentDBResponseBody) *DisableDBClusterDocumentDBResponse {
	s.Body = v
	return s
}

func (s *DisableDBClusterDocumentDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
