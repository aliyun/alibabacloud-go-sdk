// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoLabelClassificationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVideoLabelClassificationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVideoLabelClassificationResultResponse
	GetStatusCode() *int32
	SetBody(v *GetVideoLabelClassificationResultResponseBody) *GetVideoLabelClassificationResultResponse
	GetBody() *GetVideoLabelClassificationResultResponseBody
}

type GetVideoLabelClassificationResultResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoLabelClassificationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoLabelClassificationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVideoLabelClassificationResultResponse) GoString() string {
	return s.String()
}

func (s *GetVideoLabelClassificationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVideoLabelClassificationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVideoLabelClassificationResultResponse) GetBody() *GetVideoLabelClassificationResultResponseBody {
	return s.Body
}

func (s *GetVideoLabelClassificationResultResponse) SetHeaders(v map[string]*string) *GetVideoLabelClassificationResultResponse {
	s.Headers = v
	return s
}

func (s *GetVideoLabelClassificationResultResponse) SetStatusCode(v int32) *GetVideoLabelClassificationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoLabelClassificationResultResponse) SetBody(v *GetVideoLabelClassificationResultResponseBody) *GetVideoLabelClassificationResultResponse {
	s.Body = v
	return s
}

func (s *GetVideoLabelClassificationResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
