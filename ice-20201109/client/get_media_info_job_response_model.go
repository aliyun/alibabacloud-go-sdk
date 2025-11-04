// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaInfoJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMediaInfoJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMediaInfoJobResponse
	GetStatusCode() *int32
	SetBody(v *GetMediaInfoJobResponseBody) *GetMediaInfoJobResponse
	GetBody() *GetMediaInfoJobResponseBody
}

type GetMediaInfoJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMediaInfoJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMediaInfoJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMediaInfoJobResponse) GoString() string {
	return s.String()
}

func (s *GetMediaInfoJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMediaInfoJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMediaInfoJobResponse) GetBody() *GetMediaInfoJobResponseBody {
	return s.Body
}

func (s *GetMediaInfoJobResponse) SetHeaders(v map[string]*string) *GetMediaInfoJobResponse {
	s.Headers = v
	return s
}

func (s *GetMediaInfoJobResponse) SetStatusCode(v int32) *GetMediaInfoJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMediaInfoJobResponse) SetBody(v *GetMediaInfoJobResponseBody) *GetMediaInfoJobResponse {
	s.Body = v
	return s
}

func (s *GetMediaInfoJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
