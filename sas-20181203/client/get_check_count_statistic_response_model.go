// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckCountStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCheckCountStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCheckCountStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetCheckCountStatisticResponseBody) *GetCheckCountStatisticResponse
	GetBody() *GetCheckCountStatisticResponseBody
}

type GetCheckCountStatisticResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCheckCountStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCheckCountStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCheckCountStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetCheckCountStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCheckCountStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCheckCountStatisticResponse) GetBody() *GetCheckCountStatisticResponseBody {
	return s.Body
}

func (s *GetCheckCountStatisticResponse) SetHeaders(v map[string]*string) *GetCheckCountStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetCheckCountStatisticResponse) SetStatusCode(v int32) *GetCheckCountStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCheckCountStatisticResponse) SetBody(v *GetCheckCountStatisticResponseBody) *GetCheckCountStatisticResponse {
	s.Body = v
	return s
}

func (s *GetCheckCountStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
