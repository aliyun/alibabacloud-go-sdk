// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrackListByMailFromAndTagNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrackListByMailFromAndTagNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrackListByMailFromAndTagNameResponse
	GetStatusCode() *int32
	SetBody(v *GetTrackListByMailFromAndTagNameResponseBody) *GetTrackListByMailFromAndTagNameResponse
	GetBody() *GetTrackListByMailFromAndTagNameResponseBody
}

type GetTrackListByMailFromAndTagNameResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrackListByMailFromAndTagNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrackListByMailFromAndTagNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrackListByMailFromAndTagNameResponse) GoString() string {
	return s.String()
}

func (s *GetTrackListByMailFromAndTagNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrackListByMailFromAndTagNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrackListByMailFromAndTagNameResponse) GetBody() *GetTrackListByMailFromAndTagNameResponseBody {
	return s.Body
}

func (s *GetTrackListByMailFromAndTagNameResponse) SetHeaders(v map[string]*string) *GetTrackListByMailFromAndTagNameResponse {
	s.Headers = v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponse) SetStatusCode(v int32) *GetTrackListByMailFromAndTagNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponse) SetBody(v *GetTrackListByMailFromAndTagNameResponseBody) *GetTrackListByMailFromAndTagNameResponse {
	s.Body = v
	return s
}

func (s *GetTrackListByMailFromAndTagNameResponse) Validate() error {
	return dara.Validate(s)
}
