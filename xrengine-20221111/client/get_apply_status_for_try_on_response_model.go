// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplyStatusForTryOnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApplyStatusForTryOnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApplyStatusForTryOnResponse
	GetStatusCode() *int32
	SetBody(v *GetApplyStatusForTryOnResponseBody) *GetApplyStatusForTryOnResponse
	GetBody() *GetApplyStatusForTryOnResponseBody
}

type GetApplyStatusForTryOnResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplyStatusForTryOnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplyStatusForTryOnResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApplyStatusForTryOnResponse) GoString() string {
	return s.String()
}

func (s *GetApplyStatusForTryOnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApplyStatusForTryOnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApplyStatusForTryOnResponse) GetBody() *GetApplyStatusForTryOnResponseBody {
	return s.Body
}

func (s *GetApplyStatusForTryOnResponse) SetHeaders(v map[string]*string) *GetApplyStatusForTryOnResponse {
	s.Headers = v
	return s
}

func (s *GetApplyStatusForTryOnResponse) SetStatusCode(v int32) *GetApplyStatusForTryOnResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplyStatusForTryOnResponse) SetBody(v *GetApplyStatusForTryOnResponseBody) *GetApplyStatusForTryOnResponse {
	s.Body = v
	return s
}

func (s *GetApplyStatusForTryOnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
