// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAppConfigAndSaveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutAppConfigAndSaveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutAppConfigAndSaveResponse
	GetStatusCode() *int32
	SetBody(v *PutAppConfigAndSaveResponseBody) *PutAppConfigAndSaveResponse
	GetBody() *PutAppConfigAndSaveResponseBody
}

type PutAppConfigAndSaveResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutAppConfigAndSaveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutAppConfigAndSaveResponse) String() string {
	return dara.Prettify(s)
}

func (s PutAppConfigAndSaveResponse) GoString() string {
	return s.String()
}

func (s *PutAppConfigAndSaveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutAppConfigAndSaveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutAppConfigAndSaveResponse) GetBody() *PutAppConfigAndSaveResponseBody {
	return s.Body
}

func (s *PutAppConfigAndSaveResponse) SetHeaders(v map[string]*string) *PutAppConfigAndSaveResponse {
	s.Headers = v
	return s
}

func (s *PutAppConfigAndSaveResponse) SetStatusCode(v int32) *PutAppConfigAndSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *PutAppConfigAndSaveResponse) SetBody(v *PutAppConfigAndSaveResponseBody) *PutAppConfigAndSaveResponse {
	s.Body = v
	return s
}

func (s *PutAppConfigAndSaveResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
