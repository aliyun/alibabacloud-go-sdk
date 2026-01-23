// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityIdentifyResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecurityIdentifyResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecurityIdentifyResultResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecurityIdentifyResultResponseBody) *CreateSecurityIdentifyResultResponse
	GetBody() *CreateSecurityIdentifyResultResponseBody
}

type CreateSecurityIdentifyResultResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecurityIdentifyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecurityIdentifyResultResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityIdentifyResultResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityIdentifyResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecurityIdentifyResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecurityIdentifyResultResponse) GetBody() *CreateSecurityIdentifyResultResponseBody {
	return s.Body
}

func (s *CreateSecurityIdentifyResultResponse) SetHeaders(v map[string]*string) *CreateSecurityIdentifyResultResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityIdentifyResultResponse) SetStatusCode(v int32) *CreateSecurityIdentifyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityIdentifyResultResponse) SetBody(v *CreateSecurityIdentifyResultResponseBody) *CreateSecurityIdentifyResultResponse {
	s.Body = v
	return s
}

func (s *CreateSecurityIdentifyResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
