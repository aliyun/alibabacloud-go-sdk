// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateDesignateExecutorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateDesignateExecutorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateDesignateExecutorsResponse
	GetStatusCode() *int32
	SetBody(v *OperateDesignateExecutorsResponseBody) *OperateDesignateExecutorsResponse
	GetBody() *OperateDesignateExecutorsResponseBody
}

type OperateDesignateExecutorsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateDesignateExecutorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateDesignateExecutorsResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateDesignateExecutorsResponse) GoString() string {
	return s.String()
}

func (s *OperateDesignateExecutorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateDesignateExecutorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateDesignateExecutorsResponse) GetBody() *OperateDesignateExecutorsResponseBody {
	return s.Body
}

func (s *OperateDesignateExecutorsResponse) SetHeaders(v map[string]*string) *OperateDesignateExecutorsResponse {
	s.Headers = v
	return s
}

func (s *OperateDesignateExecutorsResponse) SetStatusCode(v int32) *OperateDesignateExecutorsResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateDesignateExecutorsResponse) SetBody(v *OperateDesignateExecutorsResponseBody) *OperateDesignateExecutorsResponse {
	s.Body = v
	return s
}

func (s *OperateDesignateExecutorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
