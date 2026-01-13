// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSplitTrafficControlTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SplitTrafficControlTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SplitTrafficControlTargetResponse
	GetStatusCode() *int32
	SetBody(v *SplitTrafficControlTargetResponseBody) *SplitTrafficControlTargetResponse
	GetBody() *SplitTrafficControlTargetResponseBody
}

type SplitTrafficControlTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SplitTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SplitTrafficControlTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s SplitTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *SplitTrafficControlTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SplitTrafficControlTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SplitTrafficControlTargetResponse) GetBody() *SplitTrafficControlTargetResponseBody {
	return s.Body
}

func (s *SplitTrafficControlTargetResponse) SetHeaders(v map[string]*string) *SplitTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *SplitTrafficControlTargetResponse) SetStatusCode(v int32) *SplitTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *SplitTrafficControlTargetResponse) SetBody(v *SplitTrafficControlTargetResponseBody) *SplitTrafficControlTargetResponse {
	s.Body = v
	return s
}

func (s *SplitTrafficControlTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
