// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterProgramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCasterProgramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCasterProgramResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCasterProgramResponseBody) *ModifyCasterProgramResponse
	GetBody() *ModifyCasterProgramResponseBody
}

type ModifyCasterProgramResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCasterProgramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCasterProgramResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterProgramResponse) GoString() string {
	return s.String()
}

func (s *ModifyCasterProgramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCasterProgramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCasterProgramResponse) GetBody() *ModifyCasterProgramResponseBody {
	return s.Body
}

func (s *ModifyCasterProgramResponse) SetHeaders(v map[string]*string) *ModifyCasterProgramResponse {
	s.Headers = v
	return s
}

func (s *ModifyCasterProgramResponse) SetStatusCode(v int32) *ModifyCasterProgramResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCasterProgramResponse) SetBody(v *ModifyCasterProgramResponseBody) *ModifyCasterProgramResponse {
	s.Body = v
	return s
}

func (s *ModifyCasterProgramResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
