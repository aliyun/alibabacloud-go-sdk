// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePbcVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePbcVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePbcVersionResponse
	GetStatusCode() *int32
	SetBody(v *PbcVersion) *UpdatePbcVersionResponse
	GetBody() *PbcVersion
}

type UpdatePbcVersionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcVersion        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePbcVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePbcVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdatePbcVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePbcVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePbcVersionResponse) GetBody() *PbcVersion {
	return s.Body
}

func (s *UpdatePbcVersionResponse) SetHeaders(v map[string]*string) *UpdatePbcVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdatePbcVersionResponse) SetStatusCode(v int32) *UpdatePbcVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePbcVersionResponse) SetBody(v *PbcVersion) *UpdatePbcVersionResponse {
	s.Body = v
	return s
}

func (s *UpdatePbcVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
