// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSynchronizeResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SynchronizeResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SynchronizeResourceResponse
	GetStatusCode() *int32
	SetBody(v *SynchronizeResourceResponseBody) *SynchronizeResourceResponse
	GetBody() *SynchronizeResourceResponseBody
}

type SynchronizeResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SynchronizeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SynchronizeResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s SynchronizeResourceResponse) GoString() string {
	return s.String()
}

func (s *SynchronizeResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SynchronizeResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SynchronizeResourceResponse) GetBody() *SynchronizeResourceResponseBody {
	return s.Body
}

func (s *SynchronizeResourceResponse) SetHeaders(v map[string]*string) *SynchronizeResourceResponse {
	s.Headers = v
	return s
}

func (s *SynchronizeResourceResponse) SetStatusCode(v int32) *SynchronizeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *SynchronizeResourceResponse) SetBody(v *SynchronizeResourceResponseBody) *SynchronizeResourceResponse {
	s.Body = v
	return s
}

func (s *SynchronizeResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
