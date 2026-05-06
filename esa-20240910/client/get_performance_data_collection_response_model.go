// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPerformanceDataCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPerformanceDataCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPerformanceDataCollectionResponse
	GetStatusCode() *int32
	SetBody(v *GetPerformanceDataCollectionResponseBody) *GetPerformanceDataCollectionResponse
	GetBody() *GetPerformanceDataCollectionResponseBody
}

type GetPerformanceDataCollectionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPerformanceDataCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPerformanceDataCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPerformanceDataCollectionResponse) GoString() string {
	return s.String()
}

func (s *GetPerformanceDataCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPerformanceDataCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPerformanceDataCollectionResponse) GetBody() *GetPerformanceDataCollectionResponseBody {
	return s.Body
}

func (s *GetPerformanceDataCollectionResponse) SetHeaders(v map[string]*string) *GetPerformanceDataCollectionResponse {
	s.Headers = v
	return s
}

func (s *GetPerformanceDataCollectionResponse) SetStatusCode(v int32) *GetPerformanceDataCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPerformanceDataCollectionResponse) SetBody(v *GetPerformanceDataCollectionResponseBody) *GetPerformanceDataCollectionResponse {
	s.Body = v
	return s
}

func (s *GetPerformanceDataCollectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
