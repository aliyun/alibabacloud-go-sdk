// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnlineSeatInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOnlineSeatInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOnlineSeatInformationResponse
	GetStatusCode() *int32
	SetBody(v *GetOnlineSeatInformationResponseBody) *GetOnlineSeatInformationResponse
	GetBody() *GetOnlineSeatInformationResponseBody
}

type GetOnlineSeatInformationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOnlineSeatInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOnlineSeatInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOnlineSeatInformationResponse) GoString() string {
	return s.String()
}

func (s *GetOnlineSeatInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOnlineSeatInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOnlineSeatInformationResponse) GetBody() *GetOnlineSeatInformationResponseBody {
	return s.Body
}

func (s *GetOnlineSeatInformationResponse) SetHeaders(v map[string]*string) *GetOnlineSeatInformationResponse {
	s.Headers = v
	return s
}

func (s *GetOnlineSeatInformationResponse) SetStatusCode(v int32) *GetOnlineSeatInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOnlineSeatInformationResponse) SetBody(v *GetOnlineSeatInformationResponseBody) *GetOnlineSeatInformationResponse {
	s.Body = v
	return s
}

func (s *GetOnlineSeatInformationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
