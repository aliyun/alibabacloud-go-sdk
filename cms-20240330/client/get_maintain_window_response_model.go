// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMaintainWindowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMaintainWindowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMaintainWindowResponse
	GetStatusCode() *int32
	SetBody(v *GetMaintainWindowResponseBody) *GetMaintainWindowResponse
	GetBody() *GetMaintainWindowResponseBody
}

type GetMaintainWindowResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMaintainWindowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMaintainWindowResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMaintainWindowResponse) GoString() string {
	return s.String()
}

func (s *GetMaintainWindowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMaintainWindowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMaintainWindowResponse) GetBody() *GetMaintainWindowResponseBody {
	return s.Body
}

func (s *GetMaintainWindowResponse) SetHeaders(v map[string]*string) *GetMaintainWindowResponse {
	s.Headers = v
	return s
}

func (s *GetMaintainWindowResponse) SetStatusCode(v int32) *GetMaintainWindowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMaintainWindowResponse) SetBody(v *GetMaintainWindowResponseBody) *GetMaintainWindowResponse {
	s.Body = v
	return s
}

func (s *GetMaintainWindowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
