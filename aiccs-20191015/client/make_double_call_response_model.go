// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakeDoubleCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MakeDoubleCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MakeDoubleCallResponse
	GetStatusCode() *int32
	SetBody(v *MakeDoubleCallResponseBody) *MakeDoubleCallResponse
	GetBody() *MakeDoubleCallResponseBody
}

type MakeDoubleCallResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MakeDoubleCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MakeDoubleCallResponse) String() string {
	return dara.Prettify(s)
}

func (s MakeDoubleCallResponse) GoString() string {
	return s.String()
}

func (s *MakeDoubleCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MakeDoubleCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MakeDoubleCallResponse) GetBody() *MakeDoubleCallResponseBody {
	return s.Body
}

func (s *MakeDoubleCallResponse) SetHeaders(v map[string]*string) *MakeDoubleCallResponse {
	s.Headers = v
	return s
}

func (s *MakeDoubleCallResponse) SetStatusCode(v int32) *MakeDoubleCallResponse {
	s.StatusCode = &v
	return s
}

func (s *MakeDoubleCallResponse) SetBody(v *MakeDoubleCallResponseBody) *MakeDoubleCallResponse {
	s.Body = v
	return s
}

func (s *MakeDoubleCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
