// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAgOneKeyOnlyCheckerTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyAgOneKeyOnlyCheckerTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyAgOneKeyOnlyCheckerTaskResponse
	GetStatusCode() *int32
	SetBody(v *ApplyAgOneKeyOnlyCheckerTaskResponseBody) *ApplyAgOneKeyOnlyCheckerTaskResponse
	GetBody() *ApplyAgOneKeyOnlyCheckerTaskResponseBody
}

type ApplyAgOneKeyOnlyCheckerTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyAgOneKeyOnlyCheckerTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyAgOneKeyOnlyCheckerTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyAgOneKeyOnlyCheckerTaskResponse) GoString() string {
	return s.String()
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponse) GetBody() *ApplyAgOneKeyOnlyCheckerTaskResponseBody {
	return s.Body
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponse) SetHeaders(v map[string]*string) *ApplyAgOneKeyOnlyCheckerTaskResponse {
	s.Headers = v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponse) SetStatusCode(v int32) *ApplyAgOneKeyOnlyCheckerTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponse) SetBody(v *ApplyAgOneKeyOnlyCheckerTaskResponseBody) *ApplyAgOneKeyOnlyCheckerTaskResponse {
	s.Body = v
	return s
}

func (s *ApplyAgOneKeyOnlyCheckerTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
