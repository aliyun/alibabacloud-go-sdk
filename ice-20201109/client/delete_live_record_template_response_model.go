// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRecordTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveRecordTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveRecordTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveRecordTemplateResponseBody) *DeleteLiveRecordTemplateResponse
	GetBody() *DeleteLiveRecordTemplateResponseBody
}

type DeleteLiveRecordTemplateResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveRecordTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveRecordTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveRecordTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveRecordTemplateResponse) GetBody() *DeleteLiveRecordTemplateResponseBody {
	return s.Body
}

func (s *DeleteLiveRecordTemplateResponse) SetHeaders(v map[string]*string) *DeleteLiveRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveRecordTemplateResponse) SetStatusCode(v int32) *DeleteLiveRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveRecordTemplateResponse) SetBody(v *DeleteLiveRecordTemplateResponseBody) *DeleteLiveRecordTemplateResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveRecordTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
