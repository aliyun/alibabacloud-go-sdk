// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePbcVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePbcVersionResponse
	GetStatusCode() *int32
	SetBody(v *PbcVersion) *CreatePbcVersionResponse
	GetBody() *PbcVersion
}

type CreatePbcVersionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcVersion        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcVersionResponse) GoString() string {
	return s.String()
}

func (s *CreatePbcVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePbcVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePbcVersionResponse) GetBody() *PbcVersion {
	return s.Body
}

func (s *CreatePbcVersionResponse) SetHeaders(v map[string]*string) *CreatePbcVersionResponse {
	s.Headers = v
	return s
}

func (s *CreatePbcVersionResponse) SetStatusCode(v int32) *CreatePbcVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePbcVersionResponse) SetBody(v *PbcVersion) *CreatePbcVersionResponse {
	s.Body = v
	return s
}

func (s *CreatePbcVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
