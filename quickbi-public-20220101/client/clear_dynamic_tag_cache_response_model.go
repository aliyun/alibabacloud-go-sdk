// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearDynamicTagCacheResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearDynamicTagCacheResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearDynamicTagCacheResponse
	GetStatusCode() *int32
	SetBody(v *ClearDynamicTagCacheResponseBody) *ClearDynamicTagCacheResponse
	GetBody() *ClearDynamicTagCacheResponseBody
}

type ClearDynamicTagCacheResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearDynamicTagCacheResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearDynamicTagCacheResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearDynamicTagCacheResponse) GoString() string {
	return s.String()
}

func (s *ClearDynamicTagCacheResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearDynamicTagCacheResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearDynamicTagCacheResponse) GetBody() *ClearDynamicTagCacheResponseBody {
	return s.Body
}

func (s *ClearDynamicTagCacheResponse) SetHeaders(v map[string]*string) *ClearDynamicTagCacheResponse {
	s.Headers = v
	return s
}

func (s *ClearDynamicTagCacheResponse) SetStatusCode(v int32) *ClearDynamicTagCacheResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearDynamicTagCacheResponse) SetBody(v *ClearDynamicTagCacheResponseBody) *ClearDynamicTagCacheResponse {
	s.Body = v
	return s
}

func (s *ClearDynamicTagCacheResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
