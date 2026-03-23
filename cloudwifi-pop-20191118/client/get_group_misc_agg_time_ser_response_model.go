// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupMiscAggTimeSerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGroupMiscAggTimeSerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGroupMiscAggTimeSerResponse
	GetStatusCode() *int32
	SetBody(v *GetGroupMiscAggTimeSerResponseBody) *GetGroupMiscAggTimeSerResponse
	GetBody() *GetGroupMiscAggTimeSerResponseBody
}

type GetGroupMiscAggTimeSerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupMiscAggTimeSerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupMiscAggTimeSerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGroupMiscAggTimeSerResponse) GoString() string {
	return s.String()
}

func (s *GetGroupMiscAggTimeSerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGroupMiscAggTimeSerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGroupMiscAggTimeSerResponse) GetBody() *GetGroupMiscAggTimeSerResponseBody {
	return s.Body
}

func (s *GetGroupMiscAggTimeSerResponse) SetHeaders(v map[string]*string) *GetGroupMiscAggTimeSerResponse {
	s.Headers = v
	return s
}

func (s *GetGroupMiscAggTimeSerResponse) SetStatusCode(v int32) *GetGroupMiscAggTimeSerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupMiscAggTimeSerResponse) SetBody(v *GetGroupMiscAggTimeSerResponseBody) *GetGroupMiscAggTimeSerResponse {
	s.Body = v
	return s
}

func (s *GetGroupMiscAggTimeSerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
