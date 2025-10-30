// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStreamingDataServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyStreamingDataServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyStreamingDataServiceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyStreamingDataServiceResponseBody) *ModifyStreamingDataServiceResponse
	GetBody() *ModifyStreamingDataServiceResponseBody
}

type ModifyStreamingDataServiceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyStreamingDataServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyStreamingDataServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyStreamingDataServiceResponse) GoString() string {
	return s.String()
}

func (s *ModifyStreamingDataServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyStreamingDataServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyStreamingDataServiceResponse) GetBody() *ModifyStreamingDataServiceResponseBody {
	return s.Body
}

func (s *ModifyStreamingDataServiceResponse) SetHeaders(v map[string]*string) *ModifyStreamingDataServiceResponse {
	s.Headers = v
	return s
}

func (s *ModifyStreamingDataServiceResponse) SetStatusCode(v int32) *ModifyStreamingDataServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyStreamingDataServiceResponse) SetBody(v *ModifyStreamingDataServiceResponseBody) *ModifyStreamingDataServiceResponse {
	s.Body = v
	return s
}

func (s *ModifyStreamingDataServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
