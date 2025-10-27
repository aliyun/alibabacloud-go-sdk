// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxDesignateInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadSchedulerxDesignateInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadSchedulerxDesignateInfoResponse
	GetStatusCode() *int32
	SetBody(v *ReadSchedulerxDesignateInfoResponseBody) *ReadSchedulerxDesignateInfoResponse
	GetBody() *ReadSchedulerxDesignateInfoResponseBody
}

type ReadSchedulerxDesignateInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadSchedulerxDesignateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadSchedulerxDesignateInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateInfoResponse) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadSchedulerxDesignateInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadSchedulerxDesignateInfoResponse) GetBody() *ReadSchedulerxDesignateInfoResponseBody {
	return s.Body
}

func (s *ReadSchedulerxDesignateInfoResponse) SetHeaders(v map[string]*string) *ReadSchedulerxDesignateInfoResponse {
	s.Headers = v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponse) SetStatusCode(v int32) *ReadSchedulerxDesignateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponse) SetBody(v *ReadSchedulerxDesignateInfoResponseBody) *ReadSchedulerxDesignateInfoResponse {
	s.Body = v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
