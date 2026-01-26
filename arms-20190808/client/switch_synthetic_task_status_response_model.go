// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchSyntheticTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchSyntheticTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchSyntheticTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *SwitchSyntheticTaskStatusResponseBody) *SwitchSyntheticTaskStatusResponse
	GetBody() *SwitchSyntheticTaskStatusResponseBody
}

type SwitchSyntheticTaskStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchSyntheticTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchSyntheticTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchSyntheticTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *SwitchSyntheticTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchSyntheticTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchSyntheticTaskStatusResponse) GetBody() *SwitchSyntheticTaskStatusResponseBody {
	return s.Body
}

func (s *SwitchSyntheticTaskStatusResponse) SetHeaders(v map[string]*string) *SwitchSyntheticTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *SwitchSyntheticTaskStatusResponse) SetStatusCode(v int32) *SwitchSyntheticTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchSyntheticTaskStatusResponse) SetBody(v *SwitchSyntheticTaskStatusResponseBody) *SwitchSyntheticTaskStatusResponse {
	s.Body = v
	return s
}

func (s *SwitchSyntheticTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
