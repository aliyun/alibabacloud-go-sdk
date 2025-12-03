// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHBaseSlbServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHBaseSlbServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHBaseSlbServerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHBaseSlbServerResponseBody) *DeleteHBaseSlbServerResponse
	GetBody() *DeleteHBaseSlbServerResponseBody
}

type DeleteHBaseSlbServerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHBaseSlbServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHBaseSlbServerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHBaseSlbServerResponse) GoString() string {
	return s.String()
}

func (s *DeleteHBaseSlbServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHBaseSlbServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHBaseSlbServerResponse) GetBody() *DeleteHBaseSlbServerResponseBody {
	return s.Body
}

func (s *DeleteHBaseSlbServerResponse) SetHeaders(v map[string]*string) *DeleteHBaseSlbServerResponse {
	s.Headers = v
	return s
}

func (s *DeleteHBaseSlbServerResponse) SetStatusCode(v int32) *DeleteHBaseSlbServerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHBaseSlbServerResponse) SetBody(v *DeleteHBaseSlbServerResponseBody) *DeleteHBaseSlbServerResponse {
	s.Body = v
	return s
}

func (s *DeleteHBaseSlbServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
