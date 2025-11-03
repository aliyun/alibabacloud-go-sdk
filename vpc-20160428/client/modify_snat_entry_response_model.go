// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnatEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySnatEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySnatEntryResponse
	GetStatusCode() *int32
	SetBody(v *ModifySnatEntryResponseBody) *ModifySnatEntryResponse
	GetBody() *ModifySnatEntryResponseBody
}

type ModifySnatEntryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySnatEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySnatEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySnatEntryResponse) GoString() string {
	return s.String()
}

func (s *ModifySnatEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySnatEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySnatEntryResponse) GetBody() *ModifySnatEntryResponseBody {
	return s.Body
}

func (s *ModifySnatEntryResponse) SetHeaders(v map[string]*string) *ModifySnatEntryResponse {
	s.Headers = v
	return s
}

func (s *ModifySnatEntryResponse) SetStatusCode(v int32) *ModifySnatEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySnatEntryResponse) SetBody(v *ModifySnatEntryResponseBody) *ModifySnatEntryResponse {
	s.Body = v
	return s
}

func (s *ModifySnatEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
