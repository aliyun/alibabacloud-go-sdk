// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDialogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDialogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDialogResponse
	GetStatusCode() *int32
	SetBody(v *CreateDialogResponseBody) *CreateDialogResponse
	GetBody() *CreateDialogResponseBody
}

type CreateDialogResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDialogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDialogResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDialogResponse) GoString() string {
	return s.String()
}

func (s *CreateDialogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDialogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDialogResponse) GetBody() *CreateDialogResponseBody {
	return s.Body
}

func (s *CreateDialogResponse) SetHeaders(v map[string]*string) *CreateDialogResponse {
	s.Headers = v
	return s
}

func (s *CreateDialogResponse) SetStatusCode(v int32) *CreateDialogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDialogResponse) SetBody(v *CreateDialogResponseBody) *CreateDialogResponse {
	s.Body = v
	return s
}

func (s *CreateDialogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
