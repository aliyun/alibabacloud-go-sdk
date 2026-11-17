// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBridgeTrialWebCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BridgeTrialWebCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BridgeTrialWebCallResponse
	GetStatusCode() *int32
	SetBody(v *BridgeTrialWebCallResponseBody) *BridgeTrialWebCallResponse
	GetBody() *BridgeTrialWebCallResponseBody
}

type BridgeTrialWebCallResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BridgeTrialWebCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BridgeTrialWebCallResponse) String() string {
	return dara.Prettify(s)
}

func (s BridgeTrialWebCallResponse) GoString() string {
	return s.String()
}

func (s *BridgeTrialWebCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BridgeTrialWebCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BridgeTrialWebCallResponse) GetBody() *BridgeTrialWebCallResponseBody {
	return s.Body
}

func (s *BridgeTrialWebCallResponse) SetHeaders(v map[string]*string) *BridgeTrialWebCallResponse {
	s.Headers = v
	return s
}

func (s *BridgeTrialWebCallResponse) SetStatusCode(v int32) *BridgeTrialWebCallResponse {
	s.StatusCode = &v
	return s
}

func (s *BridgeTrialWebCallResponse) SetBody(v *BridgeTrialWebCallResponseBody) *BridgeTrialWebCallResponse {
	s.Body = v
	return s
}

func (s *BridgeTrialWebCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
