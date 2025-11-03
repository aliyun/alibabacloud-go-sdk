// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFullNatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFullNatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFullNatEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateFullNatEntryResponseBody) *CreateFullNatEntryResponse
	GetBody() *CreateFullNatEntryResponseBody
}

type CreateFullNatEntryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFullNatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFullNatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFullNatEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateFullNatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFullNatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFullNatEntryResponse) GetBody() *CreateFullNatEntryResponseBody {
	return s.Body
}

func (s *CreateFullNatEntryResponse) SetHeaders(v map[string]*string) *CreateFullNatEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateFullNatEntryResponse) SetStatusCode(v int32) *CreateFullNatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFullNatEntryResponse) SetBody(v *CreateFullNatEntryResponseBody) *CreateFullNatEntryResponse {
	s.Body = v
	return s
}

func (s *CreateFullNatEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
