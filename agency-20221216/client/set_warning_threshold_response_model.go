// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWarningThresholdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetWarningThresholdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetWarningThresholdResponse
	GetStatusCode() *int32
	SetBody(v *SetWarningThresholdResponseBody) *SetWarningThresholdResponse
	GetBody() *SetWarningThresholdResponseBody
}

type SetWarningThresholdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWarningThresholdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWarningThresholdResponse) String() string {
	return dara.Prettify(s)
}

func (s SetWarningThresholdResponse) GoString() string {
	return s.String()
}

func (s *SetWarningThresholdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetWarningThresholdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetWarningThresholdResponse) GetBody() *SetWarningThresholdResponseBody {
	return s.Body
}

func (s *SetWarningThresholdResponse) SetHeaders(v map[string]*string) *SetWarningThresholdResponse {
	s.Headers = v
	return s
}

func (s *SetWarningThresholdResponse) SetStatusCode(v int32) *SetWarningThresholdResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWarningThresholdResponse) SetBody(v *SetWarningThresholdResponseBody) *SetWarningThresholdResponse {
	s.Body = v
	return s
}

func (s *SetWarningThresholdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
