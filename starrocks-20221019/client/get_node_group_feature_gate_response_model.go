// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeGroupFeatureGateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeGroupFeatureGateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeGroupFeatureGateResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeGroupFeatureGateResponseBody) *GetNodeGroupFeatureGateResponse
	GetBody() *GetNodeGroupFeatureGateResponseBody
}

type GetNodeGroupFeatureGateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeGroupFeatureGateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeGroupFeatureGateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeGroupFeatureGateResponse) GoString() string {
	return s.String()
}

func (s *GetNodeGroupFeatureGateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeGroupFeatureGateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeGroupFeatureGateResponse) GetBody() *GetNodeGroupFeatureGateResponseBody {
	return s.Body
}

func (s *GetNodeGroupFeatureGateResponse) SetHeaders(v map[string]*string) *GetNodeGroupFeatureGateResponse {
	s.Headers = v
	return s
}

func (s *GetNodeGroupFeatureGateResponse) SetStatusCode(v int32) *GetNodeGroupFeatureGateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeGroupFeatureGateResponse) SetBody(v *GetNodeGroupFeatureGateResponseBody) *GetNodeGroupFeatureGateResponse {
	s.Body = v
	return s
}

func (s *GetNodeGroupFeatureGateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
