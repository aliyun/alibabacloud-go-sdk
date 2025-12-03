// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatasetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDatasetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDatasetResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDatasetResponseBody) *ModifyDatasetResponse
	GetBody() *ModifyDatasetResponseBody
}

type ModifyDatasetResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatasetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatasetResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatasetResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatasetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDatasetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDatasetResponse) GetBody() *ModifyDatasetResponseBody {
	return s.Body
}

func (s *ModifyDatasetResponse) SetHeaders(v map[string]*string) *ModifyDatasetResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatasetResponse) SetStatusCode(v int32) *ModifyDatasetResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatasetResponse) SetBody(v *ModifyDatasetResponseBody) *ModifyDatasetResponse {
	s.Body = v
	return s
}

func (s *ModifyDatasetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
