// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPbcVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPbcVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPbcVersionResponse
	GetStatusCode() *int32
	SetBody(v *PbcVersion) *GetPbcVersionResponse
	GetBody() *PbcVersion
}

type GetPbcVersionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PbcVersion        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPbcVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPbcVersionResponse) GoString() string {
	return s.String()
}

func (s *GetPbcVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPbcVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPbcVersionResponse) GetBody() *PbcVersion {
	return s.Body
}

func (s *GetPbcVersionResponse) SetHeaders(v map[string]*string) *GetPbcVersionResponse {
	s.Headers = v
	return s
}

func (s *GetPbcVersionResponse) SetStatusCode(v int32) *GetPbcVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPbcVersionResponse) SetBody(v *PbcVersion) *GetPbcVersionResponse {
	s.Body = v
	return s
}

func (s *GetPbcVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
