// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStreamingDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStreamingDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStreamingDataSourceResponseBody) *ModifyStreamingDataSourceResponse
	GetBody() *ModifyStreamingDataSourceResponseBody
}

type ModifyStreamingDataSourceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStreamingDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStreamingDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyStreamingDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStreamingDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStreamingDataSourceResponse) GetBody() *ModifyStreamingDataSourceResponseBody {
	return s.Body
}

func (s *ModifyStreamingDataSourceResponse) SetHeaders(v map[string]*string) *ModifyStreamingDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyStreamingDataSourceResponse) SetStatusCode(v int32) *ModifyStreamingDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStreamingDataSourceResponse) SetBody(v *ModifyStreamingDataSourceResponseBody) *ModifyStreamingDataSourceResponse {
	s.Body = v
	return s
}

func (s *ModifyStreamingDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
