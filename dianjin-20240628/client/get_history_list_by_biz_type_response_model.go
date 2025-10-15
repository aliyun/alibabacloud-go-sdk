// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHistoryListByBizTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHistoryListByBizTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHistoryListByBizTypeResponse
	GetStatusCode() *int32
	SetBody(v *GetHistoryListByBizTypeResponseBody) *GetHistoryListByBizTypeResponse
	GetBody() *GetHistoryListByBizTypeResponseBody
}

type GetHistoryListByBizTypeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHistoryListByBizTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHistoryListByBizTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHistoryListByBizTypeResponse) GoString() string {
	return s.String()
}

func (s *GetHistoryListByBizTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHistoryListByBizTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHistoryListByBizTypeResponse) GetBody() *GetHistoryListByBizTypeResponseBody {
	return s.Body
}

func (s *GetHistoryListByBizTypeResponse) SetHeaders(v map[string]*string) *GetHistoryListByBizTypeResponse {
	s.Headers = v
	return s
}

func (s *GetHistoryListByBizTypeResponse) SetStatusCode(v int32) *GetHistoryListByBizTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHistoryListByBizTypeResponse) SetBody(v *GetHistoryListByBizTypeResponseBody) *GetHistoryListByBizTypeResponse {
	s.Body = v
	return s
}

func (s *GetHistoryListByBizTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
