// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliDingMinutesContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAliDingMinutesContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAliDingMinutesContentResponse
	GetStatusCode() *int32
	SetBody(v *GetAliDingMinutesContentResponseBody) *GetAliDingMinutesContentResponse
	GetBody() *GetAliDingMinutesContentResponseBody
}

type GetAliDingMinutesContentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAliDingMinutesContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAliDingMinutesContentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAliDingMinutesContentResponse) GoString() string {
	return s.String()
}

func (s *GetAliDingMinutesContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAliDingMinutesContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAliDingMinutesContentResponse) GetBody() *GetAliDingMinutesContentResponseBody {
	return s.Body
}

func (s *GetAliDingMinutesContentResponse) SetHeaders(v map[string]*string) *GetAliDingMinutesContentResponse {
	s.Headers = v
	return s
}

func (s *GetAliDingMinutesContentResponse) SetStatusCode(v int32) *GetAliDingMinutesContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAliDingMinutesContentResponse) SetBody(v *GetAliDingMinutesContentResponseBody) *GetAliDingMinutesContentResponse {
	s.Body = v
	return s
}

func (s *GetAliDingMinutesContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
