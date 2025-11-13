// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleSimilarMaliciousFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HandleSimilarMaliciousFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HandleSimilarMaliciousFilesResponse
	GetStatusCode() *int32
	SetBody(v *HandleSimilarMaliciousFilesResponseBody) *HandleSimilarMaliciousFilesResponse
	GetBody() *HandleSimilarMaliciousFilesResponseBody
}

type HandleSimilarMaliciousFilesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HandleSimilarMaliciousFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HandleSimilarMaliciousFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s HandleSimilarMaliciousFilesResponse) GoString() string {
	return s.String()
}

func (s *HandleSimilarMaliciousFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HandleSimilarMaliciousFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HandleSimilarMaliciousFilesResponse) GetBody() *HandleSimilarMaliciousFilesResponseBody {
	return s.Body
}

func (s *HandleSimilarMaliciousFilesResponse) SetHeaders(v map[string]*string) *HandleSimilarMaliciousFilesResponse {
	s.Headers = v
	return s
}

func (s *HandleSimilarMaliciousFilesResponse) SetStatusCode(v int32) *HandleSimilarMaliciousFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *HandleSimilarMaliciousFilesResponse) SetBody(v *HandleSimilarMaliciousFilesResponseBody) *HandleSimilarMaliciousFilesResponse {
	s.Body = v
	return s
}

func (s *HandleSimilarMaliciousFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
