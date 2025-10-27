// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxDesignateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadSchedulerxDesignateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadSchedulerxDesignateDetailResponse
	GetStatusCode() *int32
	SetBody(v *ReadSchedulerxDesignateDetailResponseBody) *ReadSchedulerxDesignateDetailResponse
	GetBody() *ReadSchedulerxDesignateDetailResponseBody
}

type ReadSchedulerxDesignateDetailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadSchedulerxDesignateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadSchedulerxDesignateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateDetailResponse) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadSchedulerxDesignateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadSchedulerxDesignateDetailResponse) GetBody() *ReadSchedulerxDesignateDetailResponseBody {
	return s.Body
}

func (s *ReadSchedulerxDesignateDetailResponse) SetHeaders(v map[string]*string) *ReadSchedulerxDesignateDetailResponse {
	s.Headers = v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponse) SetStatusCode(v int32) *ReadSchedulerxDesignateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponse) SetBody(v *ReadSchedulerxDesignateDetailResponseBody) *ReadSchedulerxDesignateDetailResponse {
	s.Body = v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
