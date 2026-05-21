// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLastUpgradeRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLastUpgradeRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLastUpgradeRecordResponse
	GetStatusCode() *int32
	SetBody(v *GetLastUpgradeRecordResponseBody) *GetLastUpgradeRecordResponse
	GetBody() *GetLastUpgradeRecordResponseBody
}

type GetLastUpgradeRecordResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLastUpgradeRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLastUpgradeRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLastUpgradeRecordResponse) GoString() string {
	return s.String()
}

func (s *GetLastUpgradeRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLastUpgradeRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLastUpgradeRecordResponse) GetBody() *GetLastUpgradeRecordResponseBody {
	return s.Body
}

func (s *GetLastUpgradeRecordResponse) SetHeaders(v map[string]*string) *GetLastUpgradeRecordResponse {
	s.Headers = v
	return s
}

func (s *GetLastUpgradeRecordResponse) SetStatusCode(v int32) *GetLastUpgradeRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLastUpgradeRecordResponse) SetBody(v *GetLastUpgradeRecordResponseBody) *GetLastUpgradeRecordResponse {
	s.Body = v
	return s
}

func (s *GetLastUpgradeRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
