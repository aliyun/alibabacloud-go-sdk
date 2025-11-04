// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSyncMediaInfoJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSyncMediaInfoJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSyncMediaInfoJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSyncMediaInfoJobResponseBody) *SubmitSyncMediaInfoJobResponse
	GetBody() *SubmitSyncMediaInfoJobResponseBody
}

type SubmitSyncMediaInfoJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSyncMediaInfoJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSyncMediaInfoJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSyncMediaInfoJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSyncMediaInfoJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSyncMediaInfoJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSyncMediaInfoJobResponse) GetBody() *SubmitSyncMediaInfoJobResponseBody {
	return s.Body
}

func (s *SubmitSyncMediaInfoJobResponse) SetHeaders(v map[string]*string) *SubmitSyncMediaInfoJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSyncMediaInfoJobResponse) SetStatusCode(v int32) *SubmitSyncMediaInfoJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSyncMediaInfoJobResponse) SetBody(v *SubmitSyncMediaInfoJobResponseBody) *SubmitSyncMediaInfoJobResponse {
	s.Body = v
	return s
}

func (s *SubmitSyncMediaInfoJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
