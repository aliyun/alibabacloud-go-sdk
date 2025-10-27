// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCanTrySasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCanTrySasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCanTrySasResponse
	GetStatusCode() *int32
	SetBody(v *GetCanTrySasResponseBody) *GetCanTrySasResponse
	GetBody() *GetCanTrySasResponseBody
}

type GetCanTrySasResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCanTrySasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCanTrySasResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCanTrySasResponse) GoString() string {
	return s.String()
}

func (s *GetCanTrySasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCanTrySasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCanTrySasResponse) GetBody() *GetCanTrySasResponseBody {
	return s.Body
}

func (s *GetCanTrySasResponse) SetHeaders(v map[string]*string) *GetCanTrySasResponse {
	s.Headers = v
	return s
}

func (s *GetCanTrySasResponse) SetStatusCode(v int32) *GetCanTrySasResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCanTrySasResponse) SetBody(v *GetCanTrySasResponseBody) *GetCanTrySasResponse {
	s.Body = v
	return s
}

func (s *GetCanTrySasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
