// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetShareResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetShareResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetShareResponse
	GetStatusCode() *int32
	SetBody(v *Share) *GetShareResponse
	GetBody() *Share
}

type GetShareResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Share             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetShareResponse) String() string {
	return dara.Prettify(s)
}

func (s GetShareResponse) GoString() string {
	return s.String()
}

func (s *GetShareResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetShareResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetShareResponse) GetBody() *Share {
	return s.Body
}

func (s *GetShareResponse) SetHeaders(v map[string]*string) *GetShareResponse {
	s.Headers = v
	return s
}

func (s *GetShareResponse) SetStatusCode(v int32) *GetShareResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShareResponse) SetBody(v *Share) *GetShareResponse {
	s.Body = v
	return s
}

func (s *GetShareResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
