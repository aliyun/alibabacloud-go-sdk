// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySyncJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySyncJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySyncJobResponse
	GetStatusCode() *int32
	SetBody(v *ModifySyncJobResponseBody) *ModifySyncJobResponse
	GetBody() *ModifySyncJobResponseBody
}

type ModifySyncJobResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySyncJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySyncJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySyncJobResponse) GoString() string {
	return s.String()
}

func (s *ModifySyncJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySyncJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySyncJobResponse) GetBody() *ModifySyncJobResponseBody {
	return s.Body
}

func (s *ModifySyncJobResponse) SetHeaders(v map[string]*string) *ModifySyncJobResponse {
	s.Headers = v
	return s
}

func (s *ModifySyncJobResponse) SetStatusCode(v int32) *ModifySyncJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySyncJobResponse) SetBody(v *ModifySyncJobResponseBody) *ModifySyncJobResponse {
	s.Body = v
	return s
}

func (s *ModifySyncJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
