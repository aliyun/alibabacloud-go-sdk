// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportFpShotJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportFpShotJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportFpShotJobResponse
	GetStatusCode() *int32
	SetBody(v *ImportFpShotJobResponseBody) *ImportFpShotJobResponse
	GetBody() *ImportFpShotJobResponseBody
}

type ImportFpShotJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportFpShotJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportFpShotJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportFpShotJobResponse) GoString() string {
	return s.String()
}

func (s *ImportFpShotJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportFpShotJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportFpShotJobResponse) GetBody() *ImportFpShotJobResponseBody {
	return s.Body
}

func (s *ImportFpShotJobResponse) SetHeaders(v map[string]*string) *ImportFpShotJobResponse {
	s.Headers = v
	return s
}

func (s *ImportFpShotJobResponse) SetStatusCode(v int32) *ImportFpShotJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportFpShotJobResponse) SetBody(v *ImportFpShotJobResponseBody) *ImportFpShotJobResponse {
	s.Body = v
	return s
}

func (s *ImportFpShotJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
