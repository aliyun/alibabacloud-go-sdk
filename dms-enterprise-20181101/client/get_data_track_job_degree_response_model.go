// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrackJobDegreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataTrackJobDegreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataTrackJobDegreeResponse
	GetStatusCode() *int32
	SetBody(v *GetDataTrackJobDegreeResponseBody) *GetDataTrackJobDegreeResponse
	GetBody() *GetDataTrackJobDegreeResponseBody
}

type GetDataTrackJobDegreeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataTrackJobDegreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataTrackJobDegreeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackJobDegreeResponse) GoString() string {
	return s.String()
}

func (s *GetDataTrackJobDegreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataTrackJobDegreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataTrackJobDegreeResponse) GetBody() *GetDataTrackJobDegreeResponseBody {
	return s.Body
}

func (s *GetDataTrackJobDegreeResponse) SetHeaders(v map[string]*string) *GetDataTrackJobDegreeResponse {
	s.Headers = v
	return s
}

func (s *GetDataTrackJobDegreeResponse) SetStatusCode(v int32) *GetDataTrackJobDegreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataTrackJobDegreeResponse) SetBody(v *GetDataTrackJobDegreeResponseBody) *GetDataTrackJobDegreeResponse {
	s.Body = v
	return s
}

func (s *GetDataTrackJobDegreeResponse) Validate() error {
	return dara.Validate(s)
}
