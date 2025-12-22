// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoCreationTaskResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoCreationTaskResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoCreationTaskResultResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoCreationTaskResultResponseBody) *GetVideoCreationTaskResultResponse
	GetBody() *GetVideoCreationTaskResultResponseBody
}

type GetVideoCreationTaskResultResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoCreationTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoCreationTaskResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreationTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetVideoCreationTaskResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoCreationTaskResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoCreationTaskResultResponse) GetBody() *GetVideoCreationTaskResultResponseBody {
	return s.Body
}

func (s *GetVideoCreationTaskResultResponse) SetHeaders(v map[string]*string) *GetVideoCreationTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetVideoCreationTaskResultResponse) SetStatusCode(v int32) *GetVideoCreationTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoCreationTaskResultResponse) SetBody(v *GetVideoCreationTaskResultResponseBody) *GetVideoCreationTaskResultResponse {
	s.Body = v
	return s
}

func (s *GetVideoCreationTaskResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
