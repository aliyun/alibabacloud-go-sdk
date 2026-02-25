// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateParameterSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateParameterSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateParameterSetResponse
	GetStatusCode() *int32
	SetBody(v *DissociateParameterSetResponseBody) *DissociateParameterSetResponse
	GetBody() *DissociateParameterSetResponseBody
}

type DissociateParameterSetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateParameterSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateParameterSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateParameterSetResponse) GoString() string {
	return s.String()
}

func (s *DissociateParameterSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateParameterSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateParameterSetResponse) GetBody() *DissociateParameterSetResponseBody {
	return s.Body
}

func (s *DissociateParameterSetResponse) SetHeaders(v map[string]*string) *DissociateParameterSetResponse {
	s.Headers = v
	return s
}

func (s *DissociateParameterSetResponse) SetStatusCode(v int32) *DissociateParameterSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateParameterSetResponse) SetBody(v *DissociateParameterSetResponseBody) *DissociateParameterSetResponse {
	s.Body = v
	return s
}

func (s *DissociateParameterSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
