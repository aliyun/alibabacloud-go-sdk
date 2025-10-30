// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountFactoryBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccountFactoryBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccountFactoryBaselineResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccountFactoryBaselineResponseBody) *DeleteAccountFactoryBaselineResponse
	GetBody() *DeleteAccountFactoryBaselineResponseBody
}

type DeleteAccountFactoryBaselineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccountFactoryBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccountFactoryBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountFactoryBaselineResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccountFactoryBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccountFactoryBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccountFactoryBaselineResponse) GetBody() *DeleteAccountFactoryBaselineResponseBody {
	return s.Body
}

func (s *DeleteAccountFactoryBaselineResponse) SetHeaders(v map[string]*string) *DeleteAccountFactoryBaselineResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccountFactoryBaselineResponse) SetStatusCode(v int32) *DeleteAccountFactoryBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccountFactoryBaselineResponse) SetBody(v *DeleteAccountFactoryBaselineResponseBody) *DeleteAccountFactoryBaselineResponse {
	s.Body = v
	return s
}

func (s *DeleteAccountFactoryBaselineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
