// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoIncrementUsageStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAutoIncrementUsageStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAutoIncrementUsageStatisticResponse
	GetStatusCode() *int32
	SetBody(v *GetAutoIncrementUsageStatisticResponseBody) *GetAutoIncrementUsageStatisticResponse
	GetBody() *GetAutoIncrementUsageStatisticResponseBody
}

type GetAutoIncrementUsageStatisticResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoIncrementUsageStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoIncrementUsageStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAutoIncrementUsageStatisticResponse) GoString() string {
	return s.String()
}

func (s *GetAutoIncrementUsageStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAutoIncrementUsageStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAutoIncrementUsageStatisticResponse) GetBody() *GetAutoIncrementUsageStatisticResponseBody {
	return s.Body
}

func (s *GetAutoIncrementUsageStatisticResponse) SetHeaders(v map[string]*string) *GetAutoIncrementUsageStatisticResponse {
	s.Headers = v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponse) SetStatusCode(v int32) *GetAutoIncrementUsageStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponse) SetBody(v *GetAutoIncrementUsageStatisticResponseBody) *GetAutoIncrementUsageStatisticResponse {
	s.Body = v
	return s
}

func (s *GetAutoIncrementUsageStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
