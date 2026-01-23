// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityWatchByObjectIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityWatchByObjectIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityWatchByObjectIdResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityWatchByObjectIdResponseBody) *GetQualityWatchByObjectIdResponse
	GetBody() *GetQualityWatchByObjectIdResponseBody
}

type GetQualityWatchByObjectIdResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityWatchByObjectIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityWatchByObjectIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityWatchByObjectIdResponse) GoString() string {
	return s.String()
}

func (s *GetQualityWatchByObjectIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityWatchByObjectIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityWatchByObjectIdResponse) GetBody() *GetQualityWatchByObjectIdResponseBody {
	return s.Body
}

func (s *GetQualityWatchByObjectIdResponse) SetHeaders(v map[string]*string) *GetQualityWatchByObjectIdResponse {
	s.Headers = v
	return s
}

func (s *GetQualityWatchByObjectIdResponse) SetStatusCode(v int32) *GetQualityWatchByObjectIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityWatchByObjectIdResponse) SetBody(v *GetQualityWatchByObjectIdResponseBody) *GetQualityWatchByObjectIdResponse {
	s.Body = v
	return s
}

func (s *GetQualityWatchByObjectIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
