// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaSequencesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddMediaSequencesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddMediaSequencesResponse
	GetStatusCode() *int32
	SetBody(v *AddMediaSequencesResponseBody) *AddMediaSequencesResponse
	GetBody() *AddMediaSequencesResponseBody
}

type AddMediaSequencesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMediaSequencesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMediaSequencesResponse) String() string {
	return dara.Prettify(s)
}

func (s AddMediaSequencesResponse) GoString() string {
	return s.String()
}

func (s *AddMediaSequencesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddMediaSequencesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddMediaSequencesResponse) GetBody() *AddMediaSequencesResponseBody {
	return s.Body
}

func (s *AddMediaSequencesResponse) SetHeaders(v map[string]*string) *AddMediaSequencesResponse {
	s.Headers = v
	return s
}

func (s *AddMediaSequencesResponse) SetStatusCode(v int32) *AddMediaSequencesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMediaSequencesResponse) SetBody(v *AddMediaSequencesResponseBody) *AddMediaSequencesResponse {
	s.Body = v
	return s
}

func (s *AddMediaSequencesResponse) Validate() error {
	return dara.Validate(s)
}
