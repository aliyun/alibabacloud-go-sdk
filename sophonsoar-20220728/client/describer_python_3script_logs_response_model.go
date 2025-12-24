// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescriberPython3ScriptLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescriberPython3ScriptLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescriberPython3ScriptLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescriberPython3ScriptLogsResponseBody) *DescriberPython3ScriptLogsResponse
	GetBody() *DescriberPython3ScriptLogsResponseBody
}

type DescriberPython3ScriptLogsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescriberPython3ScriptLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescriberPython3ScriptLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescriberPython3ScriptLogsResponse) GoString() string {
	return s.String()
}

func (s *DescriberPython3ScriptLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescriberPython3ScriptLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescriberPython3ScriptLogsResponse) GetBody() *DescriberPython3ScriptLogsResponseBody {
	return s.Body
}

func (s *DescriberPython3ScriptLogsResponse) SetHeaders(v map[string]*string) *DescriberPython3ScriptLogsResponse {
	s.Headers = v
	return s
}

func (s *DescriberPython3ScriptLogsResponse) SetStatusCode(v int32) *DescriberPython3ScriptLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescriberPython3ScriptLogsResponse) SetBody(v *DescriberPython3ScriptLogsResponseBody) *DescriberPython3ScriptLogsResponse {
	s.Body = v
	return s
}

func (s *DescriberPython3ScriptLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
