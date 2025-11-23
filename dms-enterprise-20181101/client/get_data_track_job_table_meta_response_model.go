// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataTrackJobTableMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataTrackJobTableMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataTrackJobTableMetaResponse
	GetStatusCode() *int32
	SetBody(v *GetDataTrackJobTableMetaResponseBody) *GetDataTrackJobTableMetaResponse
	GetBody() *GetDataTrackJobTableMetaResponseBody
}

type GetDataTrackJobTableMetaResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataTrackJobTableMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataTrackJobTableMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataTrackJobTableMetaResponse) GoString() string {
	return s.String()
}

func (s *GetDataTrackJobTableMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataTrackJobTableMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataTrackJobTableMetaResponse) GetBody() *GetDataTrackJobTableMetaResponseBody {
	return s.Body
}

func (s *GetDataTrackJobTableMetaResponse) SetHeaders(v map[string]*string) *GetDataTrackJobTableMetaResponse {
	s.Headers = v
	return s
}

func (s *GetDataTrackJobTableMetaResponse) SetStatusCode(v int32) *GetDataTrackJobTableMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataTrackJobTableMetaResponse) SetBody(v *GetDataTrackJobTableMetaResponseBody) *GetDataTrackJobTableMetaResponse {
	s.Body = v
	return s
}

func (s *GetDataTrackJobTableMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
