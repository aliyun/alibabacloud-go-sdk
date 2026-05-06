// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMergeDownloadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MergeDownloadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MergeDownloadResponse
	GetStatusCode() *int32
	SetBody(v *MergeDownloadResponseBody) *MergeDownloadResponse
	GetBody() *MergeDownloadResponseBody
}

type MergeDownloadResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MergeDownloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MergeDownloadResponse) String() string {
	return dara.Prettify(s)
}

func (s MergeDownloadResponse) GoString() string {
	return s.String()
}

func (s *MergeDownloadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MergeDownloadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MergeDownloadResponse) GetBody() *MergeDownloadResponseBody {
	return s.Body
}

func (s *MergeDownloadResponse) SetHeaders(v map[string]*string) *MergeDownloadResponse {
	s.Headers = v
	return s
}

func (s *MergeDownloadResponse) SetStatusCode(v int32) *MergeDownloadResponse {
	s.StatusCode = &v
	return s
}

func (s *MergeDownloadResponse) SetBody(v *MergeDownloadResponseBody) *MergeDownloadResponse {
	s.Body = v
	return s
}

func (s *MergeDownloadResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
